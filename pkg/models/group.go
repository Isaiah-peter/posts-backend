package models

import (
	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	Name   string `json:"nameofgroup"`
	UserId int64  `json:"user_id"`
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
	db = config.GetDB()
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&GroupUser{})
	db.AutoMigrate(&GroupMessages{})
}

func (g *Group) CreateGroup() *Group {
	db.NewRecord(g)
	db.Create(g)
	return g
}

func (g *GroupUser) AddGroup() *GroupUser {
	db.NewRecord(g)
	db.Create(g)
	return g
}

func (m *GroupMessages) GroupMessage() *GroupMessages {
	db.NewRecord(m)
	db.Create(m)
	return m
}
