package routes

import (
	"net/http"

	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var UserFollowers = func(router *mux.Router) {
	router.HandleFunc("/follower", controllers.Followers).Methods("POST")
	router.HandleFunc("/follower", controllers.GetFollower).Methods("GET")
	router.HandleFunc("follower/detail", controllers.GetUserFollowerDetails).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	}).Handler(router)

	http.ListenAndServe("Localhost:9900", handler)
}
