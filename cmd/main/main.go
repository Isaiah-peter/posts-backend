package main

import (
	"log"
	"net/http"

	"github.com/Isaiah-peter/posts-backend/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUser(r)
	routes.Followers(r)
	routes.Post(r)
	routes.Websoc(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:8000", r))
}
