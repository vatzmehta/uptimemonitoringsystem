package url

import "time"

type UrlResourceDto struct {
	Id               uint
	Url              string
	CrawlTimeout     uint64
	Frequency        uint64
	FailureThreshold uint64
	Enable           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
