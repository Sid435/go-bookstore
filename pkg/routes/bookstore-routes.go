package routes

import (
	"github.com/gorilla/mux"
	"github.com/sid/go-bookstore/pkg/controllers"
)

var RegisterBooktoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

var OnboardingRoutes = func(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
}
