package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
	}
}

func (r *Router) ApplyMiddleware(middlewares ...Middleware) {
	for _, middleware := range middlewares {
		r.Router.Use(middleware)
	}
}

func (r *Router) RegisterRoutes() {
	// Register
}

type Middleware func(http.Handler) http.Handler
