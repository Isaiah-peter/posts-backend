package routes

import (
	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var Websoc = func(router *mux.Router) {
	router.HandleFunc("/ws", controllers.ServeWs)
}
