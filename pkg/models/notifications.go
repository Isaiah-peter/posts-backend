package models

import (
	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	dbn *gorm.DB
)

type Notification struct {
	gorm.Model
	TypeId    int64  `json: "type_id"`
	Type      string `json: "type"`
	Viewed    bool   `json: "viewed"`
	ReciverId int64  `binding:"required" json:"user_id"`
}

func init() {
	config.Connect()
	dbn = config.GetDB()
	dbn.AutoMigrate(&Notification{})
}

func (n *Notification) CreateNotification() *Notification {
	dbn.NewRecord(n)
	dbn.Create(n)
	return n
}
