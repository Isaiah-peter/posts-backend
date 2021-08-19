package routes

import (
	"net/http"

	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var UserPost = func(router *mux.Router) {
	router.HandleFunc("/post/{id}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/timeline/all", controllers.Timeline).Methods("GET")
	router.HandleFunc("/timeline", controllers.GetUserpost).Methods("GET")
	router.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/like", controllers.Like).Methods("POST")
	router.HandleFunc("/post/{id}", controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/like/{id}", controllers.Dislike).Methods("DELETE")
	router.HandleFunc("/post/{id}", controllers.UpdatePost).Methods("PUT")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	}).Handler(router)

	http.ListenAndServe(":8080", handler)

}
