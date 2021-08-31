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
	newConversation models.Conversation
)

func AddFollowerToConversation(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	var conversation = &models.Conversation{}
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	utils.ParseBody(r, conversation)
	conversation.SenderId = verifiedID
	u := conversation.CreateConversation()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)

}

func GetConvOfUser(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	conv := []models.Conversation{}
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	userid := vars["id"]
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}
	if id == verifiedID {
		u := db.Where("sender_id=?", id).Or("recieve_id=?", id).Find(&conv).Value
		res, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "pkglication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
