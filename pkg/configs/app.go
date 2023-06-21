package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, error := gorm.Open("mysql", "root:root@/demo?charset=utf8&parseTime=True&loc=Local")
	if error != nil {
		panic(error)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
