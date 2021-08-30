package models

import (
	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	Name string `json:"nameofgroup"`
}

type GroupUser struct {
	gorm.Model
	GroupId int64 `json:"group_id"`
	UserId  int64 `json:"user_id"`
}

type GroupMessages struct {
	gorm.Model
	GroupId int64  `json:"group_id"`
	UserId  int64  `json:"user_id"`
	Message string `json:"message"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&GroupUser{})
	db.AutoMigrate(&GroupMessages{})
}
