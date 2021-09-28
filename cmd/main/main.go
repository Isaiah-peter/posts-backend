package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Isaiah-peter/posts-backend/pkg/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	os.Setenv("PORT", "8000")
	port := os.Getenv("PORT")
	adder := port
	if port == ""{
		adder = "8000"
	}
	r := mux.NewRouter()
	routes.RegisterUser(r)
	routes.Followers(r)
	routes.Post(r)
	routes.ConversationUser(r)
	routes.MessageUser(r)
	routes.Group(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(adder, handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}
