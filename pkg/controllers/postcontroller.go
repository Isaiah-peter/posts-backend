package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Isaiah-peter/posts-backend/pkg/models"
	"github.com/Isaiah-peter/posts-backend/pkg/utils"
	"github.com/gorilla/mux"
)

var (
	NewPost models.Post
	NewLike models.Like
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{}
	utils.ParseBody(r, post)
	token := utils.UseToken(r)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}

	post.UserID = verifiedID
	postDetail := post.CreatePost()
	res, err := json.Marshal(postDetail)
	if err != nil {
		res, _ := json.Marshal("you are not authorize")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{}
	utils.UseToken(r)
	utils.ParseBody(r, post)
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}

	postDetail, db := models.GetPostById(id)

	if post.Desc != "" {
		postDetail.Desc = post.Desc
	}
	if post.Img != "" {
		postDetail.Img = post.Img
	}

	db.Save(&postDetail)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}
	post, _ := models.GetPostById(id)
	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}
	post := models.DeletePost(id)
	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func Like(w http.ResponseWriter, r *http.Request) {
	post := &models.Like{}
	utils.ParseBody(r, post)
	token := utils.UseToken(r)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	post.UserId = verifiedID
	like := post.CreateLike()
	res, err := json.Marshal(like)
	if err != nil {
		res, _ := json.Marshal("you are not authorize")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Dislike(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}

	post := models.DeleteLike(id)
	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Timeline(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	posts := []models.Post{}
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	u := db.Where("user_id=?", verifiedID).Find(&posts).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	newuser := models.GetPost()
	res, _ := json.Marshal(newuser)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
