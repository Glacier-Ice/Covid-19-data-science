package model

import "time"

// PostData ...
type PostData struct {
	Uploader string `json:"uploader"`
	Archive
}

// Archive ...
type Archive struct {
	ID           uint   `json:"id" gorm:"primary_key" `
	AnnounceType int    `json:"announce_type"`
	City         string `json:"city"`
	Content      string `json:"content"`
	Link         string `json:"link"`
	LinksToPic   string `json:"links_to_pic"`
	Province     string `json:"province"`
	PublishDate  string `json:"publish_date"`
	PublishTime  string `json:"publish_time"`
	Title        string `json:"title"`
}

// GetLog ...
func (p PostData) GetLog(ip string) *Log {
	return &Log{
		IP:       ip,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		City:     p.City,
		Province: p.Province,
		Uploader: p.Uploader,
	}
}

// GetArchive ...
func (p PostData) GetArchive() *Archive {
	return &p.Archive
}

// IsLegal ...
func (p PostData) IsLegal() bool {
	if p.Title == "" || p.City == "" || p.Province == "" || p.Content == "" || p.Uploader == "" {
		return false
	}
	return true
}
