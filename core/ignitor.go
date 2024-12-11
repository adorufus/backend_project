package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

var thisRouter mux.Router

func StartRouter(prefix string) {
	router := mux.NewRouter().PathPrefix(prefix).Subrouter()

	thisRouter = *router
}

func notFoundHandler() {
	thisRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}

func Ignite(port string) {
	notFoundHandler()

	http.ListenAndServe(":"+port, &thisRouter)
}
