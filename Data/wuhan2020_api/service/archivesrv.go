package service

import "github.com/arrebole/wuhan2020_api/model"

// SaveArchive 创建返回0，更新返回1
func SaveArchive(archive *model.Archive) int {
	var item = find(archive.Link)
	// 不存在
	if item.ID == 0 {
		db.Create(archive)
		return 0
	}
	archive.ID = item.ID
	db.Update(&archive)
	return 1
}

// find 判断链接是否已存在
func find(link string) model.Archive {
	var archive model.Archive
	db.Where("link = ?", link).First(&archive)
	return archive
}

// GetArchivesByCity ...
func GetArchivesByCity(city string) []model.Archive {
	var result []model.Archive
	db.Where("city = ?", city).Find(&result)
	return result
}

// GetArchivesByProvince ...
func GetArchivesByProvince(province string) []model.Archive {
	var result []model.Archive
	db.Where("province = ?", province).Find(&result)
	return result
}
