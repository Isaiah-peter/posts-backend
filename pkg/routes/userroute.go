package routes

import (
	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUser = func(router *mux.Router) {
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
}
