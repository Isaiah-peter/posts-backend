package config

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:5432/social-app?sslmode=disable"
)

func Connect() {
	godotenv.Load()
	var url = os.Getenv("DATABASE_URL")
	if url == "" {
		url = dbSource
	}
	d, err := gorm.Open(dbDriver, url)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
