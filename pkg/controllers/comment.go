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
	NewComment      models.Comment
	NewNotification models.Notification
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	comment := &models.Comment{}
	utils.ParseBody(r, comment)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	comment.UserId = verifiedID
	u := comment.CreateComment()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func GetCommentById(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	comment := []models.Comment{}
	vars := mux.Vars(r)
	userid := vars["id"]
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}

	u := db.Where("post_id=?", id).Find(&comment).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	comment := []models.Comment{}
	u := db.Find(&comment).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetNotification(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	var not []models.Notification
	vars := mux.Vars(r)
	postId := vars["id"]
	u := db.Where("post_id=?", postId).Find(&not).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
