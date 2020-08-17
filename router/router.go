package router

import (
	"basic-golang-crud/middleware"

	"github.com/gorilla/mux"
)

// Router create and setup a router
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user", middleware.UpdateUser).Methods("PATCH", "OPTIONS")

	return router
}
