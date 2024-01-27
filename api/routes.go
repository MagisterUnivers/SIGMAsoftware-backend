package api

import (
	"github.com/gorilla/mux"
)

func UsersRoutes(router *mux.Router) {
	router.HandleFunc("/", CreateUser).Methods("POST")
	router.HandleFunc("/", GetUsers).Methods("GET")
	router.HandleFunc("/{id}", GetUser).Methods("GET")
	router.HandleFunc("/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/{id}", DeleteUser).Methods("DELETE")
}
