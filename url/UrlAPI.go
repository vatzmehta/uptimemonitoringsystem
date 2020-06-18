package url

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type UrlAPI struct {
	urlService UrlService
}

func UrlAPIProvider(service UrlService) UrlAPI {
	return UrlAPI{urlService: service}
}

func (u *UrlAPI) Create(c *gin.Context) {
	var urlDto UrlResourceDto
	err := c.BindJSON(&urlDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	url := u.urlService.Create(ToUrlResource(urlDto))
	c.JSON(http.StatusOK, gin.H{"URL": ToUrlResourceDTO(url)})
}

func (u *UrlAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	url := u.urlService.FindById(uint(id))
	if url == (UrlResource{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	u.urlService.Delete(url)
	c.Status(http.StatusOK)
}

func (u *UrlAPI) DeleteAll(c *gin.Context) {
	u.urlService.DeleteAll()
	c.Status(http.StatusOK)
}

func (u *UrlAPI) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	url := u.urlService.FindById(uint(id))
	if url == (UrlResource{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"URL": ToUrlResourceDTO(url)})
}

func (u *UrlAPI) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"URLs": ToUrlResourceDTOs(u.urlService.FindAll())})
}

func (u *UrlAPI) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	originalUrl := u.urlService.FindById(uint(id))
	if originalUrl == (UrlResource{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	var urlDTO UrlResourceDto
	err := c.BindJSON(&urlDTO)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if urlDTO.Url != "" {
		originalUrl.Url = urlDTO.Url
	}

	if urlDTO.FailureThreshold != 0 {
		originalUrl.FailureThreshold = urlDTO.FailureThreshold
	}

	if urlDTO.Frequency != 0 {
		originalUrl.Frequency = urlDTO.Frequency
	}

	if urlDTO.CrawlTimeout != 0 {
		originalUrl.CrawlTimeout = urlDTO.CrawlTimeout
	}

	//enable..

	u.urlService.Create(originalUrl)
	c.Status(http.StatusOK)

}

func (u *UrlAPI) Activate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	url := u.urlService.FindById(uint(id))
	if url == (UrlResource{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	var activate string
	activate = c.Param("activate")
	if strings.EqualFold(activate, "activate") {
		u.urlService.Activate(url, true)
		c.Status(http.StatusOK)
	} else if strings.EqualFold(activate, "deactivate") {
		u.urlService.Activate(url, false)
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusBadRequest)
	}

}
