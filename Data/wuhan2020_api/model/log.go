package model

// Log ...
type Log struct {
	ID       uint   `json:"id" gorm:"primary_key" `
	City     string `json:"city"`
	IP       string `json:"ip"`
	Province string `json:"province"`
	Time     string `json:"time"`
	Uploader string `json:"uploader"`
}
