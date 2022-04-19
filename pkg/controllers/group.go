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
	Group        models.Group
	GroupUser    models.GroupUser
	GroupMessage models.GroupMessages
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	token, ok := utils.UseToken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	group := &models.Group{}
	utils.ParseBody(r, group)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	group.UserId = verifiedID
	u := group.CreateGroup()
	res, _ := json.Marshal(u)
	user := &models.GroupUser{
		UserId:  verifiedID,
		GroupId: int64(u.ID),
	}
	user.AddGroup()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func AddGroupUser(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.UseToken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	user := &models.GroupUser{}
	utils.ParseBody(r, user)
	u := user.AddGroup()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func CreateGroupMessage(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.UseToken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	message := &models.GroupMessages{}
	utils.ParseBody(r, message)

	u := message.GroupMessage()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func GetGroupById(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.UseToken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	group := []models.Group{}
	vars := mux.Vars(r)
	userid := vars["id"]
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}
	u := db.Where("user_id=?", id).Find(&group).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetGroupUserjoined(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.UseToken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	group := []models.Group{}
	vars := mux.Vars(r)
	userid := vars["id"]
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}
	u := db.Where("ID=?", id).Find(&group).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetGroupUser(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.UseToken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	group := []models.GroupUser{}
	vars := mux.Vars(r)
	userid := vars["id"]
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}
	u := db.Where("user_id=?", id).Find(&group).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetGroupUserbygruoupid(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.UseToken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	group := []models.GroupUser{}
	vars := mux.Vars(r)
	userid := vars["id"]
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}
	u := db.Where("group_id=?", id).Find(&group).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetGroupMessage(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.UseToken(r)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	group := []models.GroupMessages{}
	vars := mux.Vars(r)
	userid := vars["id"]
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}
	u := db.Where("group_id=?", id).Find(&group).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
