package models

import (
	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	dbg *gorm.DB
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
	dbg = config.GetDB()
	dbg.AutoMigrate(&Group{})
	dbg.AutoMigrate(&GroupUser{})
	dbg.AutoMigrate(&GroupMessages{})
}

func (g *Group) CreateGroup() *Group {
	dbg.NewRecord(g)
	dbg.Create(g)
	return g
}

func (g *GroupUser) AddGroup() *GroupUser {
	dbg.NewRecord(g)
	dbg.Create(g)
	return g
}

func (m *GroupMessages) GroupMessage() *GroupMessages {
	dbg.NewRecord(m)
	dbg.Create(m)
	return m
}
