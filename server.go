package main

import (
	"encoding/json"
	"log"
	"main/core"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	//Handle Not Found
	api.NotFoundHandler = http.HandlerFunc(notFound)

	//handle middleware
	api.Use(currentRequestedRoute)

	firstVersion := api.PathPrefix("/v1").Subrouter()
	firstVersion.HandleFunc("/status", getHttpStatus)
	firstVersion.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	//Done temporary, don't touch yet ok
	core.Ignite(router)
}

func currentRequestedRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-auth-token") != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func getHttpStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	_ = json.NewDecoder(r.Body).Decode(&newUser)

	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}
