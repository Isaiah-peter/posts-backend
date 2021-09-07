package models

import (
	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	TypeId int64  `json: "type_id"`
	Type   string `json: "type"`
	Viewed bool   `json: "viewed"`
	reciverId int64  `binding:"required" json:"user_id"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Notification{})
}

func (n *Notification) CreateNotification() *Notification {
	db.NewRecord(n)
	db.Create(n)
	return n
}
