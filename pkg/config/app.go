package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "b0748306b8e0f7:745801d2@/heroku_abdaf9e81ac8a78?charset=utf8mb4&parseTime=True&loc=Local&reconnect=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
