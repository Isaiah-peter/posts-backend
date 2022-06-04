package models

import (
	"fmt"

	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/Isaiah-peter/posts-backend/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name           string ` binding:"required" json:"name"`
	UserName       string ` gorm:"unique" binding:"required" json:"username"`
	Email          string ` gorm:"unique" binding:"required" json:"email"`
	Password       string ` binding:"required" json:"password"`
	ProfilePicture string ` json:"profilepicture"`
	CoverPicture   string ` json:"coverpicture"`
	IsAdmin        bool   `json:"is_admin"`
	Desc           string `json:"description"`
	City           string `json:"city"`
	From           string `json:"town"`
	Relationship   string `json:"relationship"`
}

type Follow struct {
	gorm.Model
	UserID     int64 `json:"user_id"`
	FollowerID int64 `json:"follower_id"`
}

type Token struct {
	UserID  int64
	IsAdmin bool
	jwt.StandardClaims
}

func init() {
	config.Connect()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Follow{})
}

func (u *User) CreateUser() *User {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	u.Password = hashPassword
	db.NewRecord(u)
	db.Create(u)
	return u
}

func GetUser() []User {
	var User []User
	db.Find(User)
	return User
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user
}

func (f *Follow) CreateFollower() *Follow {
	db.NewRecord(f)
	db.Create(f)
	return f
}

func Deletefollower(Id int64) Follow {
	var user Follow
	db.Where("follower_id=?", Id).Delete(user)
	return user
}
