package web

import (
	"go-react/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	var r = mux.NewRouter()

	routes.StartRouter(r)

	http.ListenAndServe(":8081", r)
}
