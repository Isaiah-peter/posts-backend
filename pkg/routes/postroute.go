package routes

import (
	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var Post = func(router *mux.Router) {
	router.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/like", controllers.Like).Methods("POST")
	router.HandleFunc("/like", controllers.GetLike).Methods("GET")
	router.HandleFunc("/timeline/all", controllers.Timeline).Methods("GET")
	router.HandleFunc("/timeline", controllers.GetUserpost).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/like/{id}", controllers.Dislike).Methods("DELETE")

}
