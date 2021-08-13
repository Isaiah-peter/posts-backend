package routes

import (
	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var UserFollowers = func(router *mux.Router) {
	router.HandleFunc("/follower", controllers.Followers).Methods("POST")
}
