package routes

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewRest() *mux.Router {
	r := mux.NewRouter()
	RegisterUser(r)
	UserFollowers(r)
	UserPost(r)
	return r
}
