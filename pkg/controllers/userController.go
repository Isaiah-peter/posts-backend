package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Isaiah-peter/posts-backend/pkg/config"
	"github.com/Isaiah-peter/posts-backend/pkg/models"
	"github.com/Isaiah-peter/posts-backend/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var (
	db      = config.GetDB()
	NewUser models.User
)

func Register(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	u := newUser.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	err := json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	u := FindOne(newUser.Email, newUser.Password)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(u)
}

func FindOne(email string, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp

	}

	expireAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID:  int64(user.ID),
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	token := utils.UseToken(r)
	utils.ParseBody(r, user)
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}

	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}

	if verifiedID == id {
		userDetail, dr := models.GetUserById(id)

		if user.Password != "" {
			hashpassword, err := utils.HashPassword(user.Password)
			if err != nil {
				panic(err)
			}
			userDetail.Password = hashpassword
		}
		if user.UserName != "" {
			userDetail.UserName = user.UserName
			fmt.Println("userDetailfmt:", userDetail.UserName)
		}
		if user.Email != "" {
			userDetail.Email = user.Email
		}
		if user.Desc != "" {
			userDetail.Desc = user.Desc
		}
		if user.Relationship != "" {
			userDetail.Relationship = user.Relationship
		}
		if user.City != "" {
			userDetail.City = user.City
			fmt.Println("city: ", user.City)
		}
		if user.ProfilePicture != "" {
			userDetail.ProfilePicture = user.ProfilePicture
			fmt.Println("ProfilePicture: ", user.ProfilePicture)
		}
		if user.CoverPicture != "" {
			userDetail.CoverPicture = user.CoverPicture
			fmt.Println("CoverPicture: ", user.CoverPicture)
		}
		if user.From != "" {
			userDetail.From = user.From
			fmt.Println("From: ", user.From)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		dr.Save(&userDetail)
	}

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	vars := mux.Vars(r)
	userId := vars["id"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}

	userDetail, _ := models.GetUserById(id)
	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}

	if verifiedID == id {
		user := models.DeleteUser(id)
		res, _ := json.Marshal(user)
		w.Header().Set("Content-Type", "pkglication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	follower := []models.User{}
	u := db.Find(&follower).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
