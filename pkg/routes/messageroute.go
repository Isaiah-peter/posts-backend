package routes

import (
	"github.com/Isaiah-peter/posts-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var MessageUser = func(router *mux.Router) {
  router.HandleFunc("/welcom", controllers.Welcome).Methods("GET")
}
