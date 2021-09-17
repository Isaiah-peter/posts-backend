package routes

import (
	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var ConversationUser = func(router *mux.Router) {
	router.HandleFunc("/addfollowertomessage", controllers.AddFollowerToConversation).Methods("POST")
	router.HandleFunc("/getconversation/{id}", controllers.GetConvOfUser).Methods("GET")
	router.HandleFunc("/search", controllers.Search).Methods("GET")
}
