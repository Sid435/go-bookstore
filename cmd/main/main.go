package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sid/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()             // getting the route
	routes.RegisterBooktoreRoutes(r) // redirecting the route to routes folder

	http.Handle("/", r)

	// ListenAndServe allows us to create a server at the defined port
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
