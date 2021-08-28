package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Isaiah-peter/posts-backend/pkg/models"
	"github.com/Isaiah-peter/posts-backend/pkg/utils"
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
	res, err := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)

}
