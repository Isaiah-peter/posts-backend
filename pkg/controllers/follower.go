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

func GetFollower(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	follower := []models.Follow{}
	u := db.Find(&follower).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserFollowerDetails(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	followers := []models.Follow{}
	user := []models.User{}
	var ids []string
	vars := mux.Vars(r)
	id := vars["id"]
	verifiedID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic(err)
	}
	db.Where("user_id=?", verifiedID).Find(&followers).Pluck("follower_id", &ids)
	fmt.Println("user_id IN (" + strings.Join(ids[:], ",") + ")")
	u := db.Where("ID IN (" + strings.Join(ids[:], ",") + ")").Find(&user).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
