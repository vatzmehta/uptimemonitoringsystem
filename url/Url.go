package url

import (
	"github.com/jinzhu/gorm"
)

type UrlResource struct {
	gorm.Model
	Url              string
	CrawlTimeout     uint64
	Frequency        uint64
	FailureThreshold uint64
	Enable           bool
}
