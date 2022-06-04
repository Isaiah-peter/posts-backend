package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Isaiah-peter/posts-backend/pkg/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = ":8000"
	}
	r := mux.NewRouter()
	routes.RegisterUser(r)
	routes.Followers(r)
	routes.Post(r)
	routes.ConversationUser(r)
	routes.MessageUser(r)
	routes.Group(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}
