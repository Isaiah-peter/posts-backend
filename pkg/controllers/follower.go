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
	NewFollower models.Follow
)

func Followers(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	follower := &models.Follow{}
	utils.ParseBody(r, follower)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	follower.UserID = verifiedID
	f := follower.CreateFollower()
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
