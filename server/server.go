package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// unexported const var
const (
	port   = ":8080"
	domain = "localhost"
)

// Start func start server process and init all
// config server options
func Start() {
	router := httprouter.New()
	router.GET("/", rootHandler)
	err := http.ListenAndServe(domain+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
