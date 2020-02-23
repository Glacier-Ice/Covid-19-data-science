package service

import "github.com/arrebole/wuhan2020_api/model"

// SaveLog ...
func SaveLog(log *model.Log) {
	db.Create(log)
}

// GetLogs ...
func GetLogs(limit int) []model.Log {
	var result []model.Log
	db.Order("id desc").Limit(limit).Find(&result)
	return result
}
