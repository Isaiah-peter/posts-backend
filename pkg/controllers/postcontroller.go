package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(res)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
		msg := "you are not authorize"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(msg))
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

func GetUserpost(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	followers := []models.Follow{}
	posts := []models.Post{}
	var ids []string
	vid := mux.Vars(r)
	id := vid["id"]
	verifiedID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}
	db.Where("user_id=?", verifiedID).Find(&followers).Pluck("follower_id", &ids)
	ids = append(ids, strconv.FormatInt(verifiedID, 10))
	fmt.Println("user_id IN (" + strings.Join(ids[:], ",") + ")")
	u := db.Where("user_id IN (" + strings.Join(ids[:], ",") + ")").Preload("Tag").Preload("Comment").Preload("Like").Find(&posts).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLike(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	vars := mux.Vars(r)
	like := []models.Like{}
	ids := vars["id"]
	u := db.Where("post_id=?", ids).Find(&like).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserpostUsingName(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	var user []models.User
	var followers []models.Follow
	var posts []models.Post
	var ids []string
	var id []string
	username := r.URL.Query()["username"]

	db.Where("user_name=?", username).Find(&user).Pluck("ID", &id)
	fmt.Println("user_id IN (" + strings.Join(id[:], ",") + ")")
	db.Where("user_id IN ("+strings.Join(id[:], ",")+")").Find(&followers).Pluck("follower_id", &ids)
	ids = append(ids, strings.Join(id[:], ","))
	fmt.Println("user_id IN (" + strings.Join(ids[:], ",") + ")")
	u := db.Where("user_id IN (" + strings.Join(ids[:], ",") + ")").Preload("Tag").Preload("Comment").Preload("Like").Find(&posts).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
