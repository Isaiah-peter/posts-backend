package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "isaiah:Etanuwoma18@/social?charset=utf8mb4&parseTime=True&loc=Local")
	c, erre := gorm.Open("mysql", os.Getenv("CLEARDB_DATABASE_URL"))
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
