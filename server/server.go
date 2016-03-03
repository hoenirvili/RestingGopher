package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// unexported const var
const (
	// server settings
	port   = ":8080"
	domain = "localhost"
	// db settings
	dbName   = "@/gophercloud"
	password = "gopherinthecloud"
	username = "golang:"
)

// Start func start server process and init all
// config server options
func Start() {
	// make a new httprouter
	router := httprouter.New()
	// make just a new logger
	Logger = newLogger()
	// GET request
	router.GET("/", rootHandler)
	router.GET("/articles/:id", articlesHandler)
	router.GET("/categories/", categoriesHandler)

	// just make static public directory
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	// start the server
	err := http.ListenAndServe(domain+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
