package models

import (
	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	dbsa *gorm.DB
)

type Post struct {
	gorm.Model
	UserID int64  `binding:"required" json:"user_id"`
	Desc   string `json:"description"`
	Img    string `json:"image"`
}

type Like struct {
	gorm.Model
	UserId int64 `binding:"required" json:"user_id"`
	PostId int64 `binding:"required" json:"post_id"`
}

type Comment struct {
	gorm.Model
	Message string `binding:"required" json:"comment"`
	UserId  int64  `binding:"required" json:"user_id"`
	PostId  int64  `binding:"required" json:"post_id"`
}

type Tag struct {
	gorm.Model
	UserId int64 `binding:"required" json:"user_id"`
	PostId int64 `binding:"required" json:"post_id"`
}

func init() {
	config.Connect()
	dbsa = config.GetDB()
	dbsa.AutoMigrate(&Post{})
	dbsa.AutoMigrate(&Like{})
	dbsa.AutoMigrate(&Comment{})
	dbsa.AutoMigrate(&Tag{})
}

func (u *Post) CreatePost() *Post {
	dbsa.NewRecord(u)
	dbsa.Create(u)
	return u
}

func GetPost() []Post {
	var Post []Post
	dbsa.Find(&Post)
	return Post
}

func GetPostById(Id int64) (*Post, *gorm.DB) {
	var getPost Post
	db := dbsa.Where("ID=?", Id).Find(&getPost)
	return &getPost, db
}

func DeletePost(Id int64) Post {
	var post Post
	dbsa.Where("ID=?", Id).Delete(post)
	return post
}

func (u *Like) CreateLike() *Like {
	dbsa.NewRecord(u)
	dbsa.Create(u)
	return u
}

func DeleteLike(Id int64) Like {
	var like Like
	dbsa.Where("post_id=?", Id).Delete(like)
	return like
}

func GetAllPostById(Id int64) *Post {
	var getPost Post
	dbsa.Where("user_id=?", Id).Find(&getPost)
	return &getPost
}

func (c *Comment) CreateComment() *Comment {
	dbsa.NewRecord(c)
	dbsa.Create(c)
	return c
}

func (t *Tag)CreateTag() *Tag {
	dbsa.NewRecord(t)
	dbsa.Create(t)
	return t
}

