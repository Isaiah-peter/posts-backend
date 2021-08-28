package controllers

import (
	"encoding/json"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	res := json.RawMessage("dnvnnncdn")
	w.Write(res)
}
