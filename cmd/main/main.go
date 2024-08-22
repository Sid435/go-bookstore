package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sid/go-bookstore/pkg/controllers"
	"github.com/sid/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter() // getting the route

	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.Use(controllers.JWTAuthMiddleware)
	routes.RegisterBooktoreRoutes(bookRouter)

	onBoard := r.PathPrefix("/auth").Subrouter()
	routes.OnboardingRoutes(onBoard)

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
