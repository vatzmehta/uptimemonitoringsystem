package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gojektech/heimdall/hystrix"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strings"
	"sync"
	"time"
	"uptimemonitoringsystem/audit"
	"uptimemonitoringsystem/url"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost)/crawler?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&url.UrlResource{})
	db.AutoMigrate(&audit.Audit{})
	return db
}

func main() {
	db := initDB()
	defer db.Close()
	auditRepo := audit.RepoProvider(db)
	urlRepo := url.RepoProvider(db)
	urlService := url.UrlServiceProvider(urlRepo)
	urlAPI := url.UrlAPIProvider(urlService)

	if len(os.Args) > 1 {
		argsWithoutProg := os.Args[1]
		if strings.EqualFold(argsWithoutProg, "audit") {
			go cmdApp(urlService, auditRepo)
		}
	}

	r := gin.Default()

	r.GET("/v1/urls", urlAPI.FindAll)
	r.GET("/v1/urls/:id", urlAPI.FindById)
	r.POST("/v1/urls/create", urlAPI.Create)
	r.DELETE("/v1/urls", urlAPI.DeleteAll)
	r.DELETE("/v1/urls/:id", urlAPI.Delete)
	r.PATCH("/v1/urls/:id", urlAPI.Update)
	r.PATCH("/v1/urls/:id/:activate", urlAPI.Activate)

	err := r.Run(":8082")
	if err != nil {
		panic(err)
	}

}

func cmdApp(urlService url.UrlService, auditRepo audit.AuditRepo) {
	failureThresholdMap := make(map[string]uint64)
	activeUrls := urlService.FindAllActive()
	initializeMap(failureThresholdMap, activeUrls)
	mutex := &sync.Mutex{}
	for _, v := range activeUrls {
		incoming := make(chan url.UrlResource)
		go crawl(urlService, auditRepo, incoming, failureThresholdMap, mutex)
		incoming <- v
	}

}

func initializeMap(thresholdMap map[string]uint64, urls []url.UrlResource) {
	for _, v := range urls {
		thresholdMap[v.Url] = 0
	}
}

func crawl(urlService url.UrlService, auditRepo audit.AuditRepo, incoming chan url.UrlResource,
	failureThresholdMap map[string]uint64, lock *sync.Mutex) {
	for v := range incoming {
		lock.Lock()
		if failureThresholdMap[v.Url] <= v.FailureThreshold {
			lock.Unlock()
			client := hystrix.NewClient(
				hystrix.WithHTTPTimeout(time.Duration(v.CrawlTimeout)*time.Second*2),
				hystrix.WithHystrixTimeout(time.Duration(v.CrawlTimeout) * time.Second),
			)
			startTime := time.Now()
			resp, err := client.Get(v.Url, nil)
			endTime := time.Now()
			var statusCode int
			if resp == nil {
				statusCode = 0 //timeout (408)
			} else {
				statusCode = resp.StatusCode
			}
			if err != nil || resp == nil || resp.StatusCode != 200 {
				//failure count
				lock.Lock()
				failureThresholdMap[v.Url] += 1
				lock.Unlock()
			}

			// create a audit
			auditRepo.Save(audit.Audit{Url: v.Url, CrawlStartTime: startTime, CrawlEndTime: endTime, HttpStatus: statusCode})
			time.Sleep(time.Duration(v.Frequency) * time.Second)
		} else {
			//mark the url inactive
			urlService.Activate(v, false)
			// do not push to same channel
			continue
		}
		// push it back to same channel
		go func() {
			incoming <- v
		}()
	}
}
