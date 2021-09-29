package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "isaiah:Etanuwoma18@/social?charset=utf8mb4&parseTime=True&loc=Local")
	c, erre := gorm.Open("mysql", "ba4efbf4b7a5b2:f5e2d1c4@us-cdbr-east-04.cleardb.com/heroku_2f48e549c2f3b08?charset=utf8mb4&parseTime=True&loc=Local")
  	if err != nil {
		panic(err)
	}
	if erre != nil {
		panic(erre)
	}
	if c == nil {
		db = d
	}
	db = c
}

func GetDB() *gorm.DB {
	return db
}
