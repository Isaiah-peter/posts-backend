package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
    godotenv.Load()


	username := os.Getenv("db_user")
	if username == "" {
		username = "isaiah"
	}
	password := os.Getenv("db_pass")
	if password == "" {
		password = "Etanuwoma18"
	}
	dbName := os.Getenv("db_name")
	if dbName== "" {
		dbName = "social"
	}
	dbUri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",username ,password , dbName ) //Build connection string
	fmt.Println(dbUri)
	d, err := gorm.Open("mysql", dbUri)
  	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
