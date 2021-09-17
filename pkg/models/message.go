package models

import (
	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type Message struct {
	gorm.Model
	ConversationId int64  `json:"conversation_id"`
	Sender         int64  `json:"sender"`
	Text           string `json:"text"`
}

type Conversation struct {
	gorm.Model
	RecieveId int64 `json:"recieve_id"`
	SenderId  int64 `json:"sender_id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Conversation{})
	db.AutoMigrate(&Message{})
}



func (c *Message) CreateMessage() *Message {
	db.NewRecord(c)
	db.Create(c)
	return c
}

func (c *Conversation) CreateConversation() *Conversation {

	db.FirstOrCreate(c)
	return c
}
