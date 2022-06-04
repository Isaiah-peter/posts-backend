package models

import (
	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	UserID  int64  `binding:"required" json:"user_id"`
	Desc    string `json:"description"`
	Img     string `json:"image"`
	Tag     []Tag
	Comment []Comment
	Like    []Like
}

type Like struct {
	gorm.Model
	UserId int64 `binding:"required" json:"user_id"`
	PostId int64 `binding:"required" json:"post_id"`
}

type Comment struct {
	gorm.Model
	Mesge  string `binding:"required" json:"comment"`
	UserId int64  `binding:"required" json:"user_id"`
	PostId int64  `binding:"required" json:"post_id"`
}

type Tag struct {
	gorm.Model
	UserId int64 `binding:"required" json:"user_id"`
	PostId int64 `binding:"required" json:"post_id"`
}

func init() {
	config.Connect()
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Like{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Tag{})
}

func (u *Post) CreatePost() *Post {
	db.NewRecord(u)
	db.Create(u)
	return u
}

func GetPost() []Post {
	var Post []Post
	db.Find(&Post)
	return Post
}

func GetPostById(Id int64) (*Post, *gorm.DB) {
	var getPost Post
	db := db.Where("ID=?", Id).Preload("Tag").Preload("Comment").Preload("Like").Find(&getPost)
	return &getPost, db
}

func DeletePost(Id int64) Post {
	var post Post
	db.Where("ID=?", Id).Delete(post)
	return post
}

func (u *Like) CreateLike() *Like {
	db.NewRecord(u)
	db.Create(u)
	return u
}

func DeleteLike(Id int64) Like {
	var like Like
	db.Where("post_id=?", Id).Delete(like)
	return like
}

func GetAllPostById(Id int64) *Post {
	var getPost Post
	db.Where("user_id=?", Id).Find(&getPost)
	return &getPost
}

func (c *Comment) CreateComment() *Comment {
	db.NewRecord(c)
	db.Create(c)
	return c
}

func (t *Tag) CreateTag() *Tag {
	db.NewRecord(t)
	db.Create(t)
	return t
}
