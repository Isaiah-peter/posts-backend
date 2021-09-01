package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Isaiah-peter/posts-backend/pkg/models"
	"github.com/Isaiah-peter/posts-backend/pkg/utils"
	"github.com/gorilla/mux"
)

var (
	newMessage models.Message
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	var message = &models.Message{}

	utils.ParseBody(r, message)
	u := message.CreateMessage()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)

}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	message := []models.Message{}

	vars := mux.Vars(r)
	userid := vars["id"]
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}

	u := db.Where("conversation_id=?", id).Find(&message).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
