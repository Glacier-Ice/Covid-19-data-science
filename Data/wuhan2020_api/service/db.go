package service

import (
	"os"

	"github.com/jinzhu/gorm"

	// mysql 驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	db = connect()
}

func password() string {
	return os.Getenv("PASSWORD")
}

func connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:"+password()+"@/db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	return db
}
