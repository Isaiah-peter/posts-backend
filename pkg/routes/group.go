package routes

import (
	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var Group = func(router *mux.Router) {
	router.HandleFunc("/group", controllers.CreateGroup).Methods("POST")
	router.HandleFunc("/addfriend", controllers.AddGroupUser).Methods("POST")
	router.HandleFunc("/groupmessage", controllers.CreateGroupMessage).Methods("POST")
	router.HandleFunc("/comment", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/group", controllers.GetGroupById).Methods("GET")
	router.HandleFunc("/groupuserjoin/{id}", controllers.GetGroupUserjoined).Methods("GET")
	router.HandleFunc("/groupuser/{id}", controllers.GetGroupUser).Methods("GET")
	router.HandleFunc("/groupmessage/{id}", controllers.GetGroupMessage).Methods("GET")
	router.HandleFunc("/comment/{id}", controllers.GetCommentById).Methods("GET")
	router.HandleFunc("/comment", controllers.GetComment).Methods("GET")
}
