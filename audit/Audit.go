package audit

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Audit struct {
	gorm.Model
	Url            string
	CrawlStartTime time.Time
	CrawlEndTime   time.Time
	HttpStatus     int
}
