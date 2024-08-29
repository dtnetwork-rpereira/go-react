package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartRouter(r *mux.Router) {
	r.HandleFunc("/", HomeLander)
}

func HomeLander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá 2"))
}
