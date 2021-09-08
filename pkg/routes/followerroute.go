package routes

import (
	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var Followers = func(router *mux.Router) {
	router.HandleFunc("/follower", controllers.Followers).Methods("POST")
	router.HandleFunc("/tag", controllers.CreateTag).Methods("POST")
	router.HandleFunc("/follower", controllers.GetFollower).Methods("GET")
	router.HandleFunc("/follower/{id}", controllers.GetUserFollowerDetails).Methods("GET")
	router.HandleFunc("/tag/{id}", controllers.GetTag).Methods("GET")
	router.HandleFunc("/follower/{id}", controllers.Unfollow).Methods("DELETE")
}
