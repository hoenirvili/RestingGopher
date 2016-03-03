// Copyright [2016] [hoenir]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

type customMethodNotAllowed struct {
}
type customNotFound struct {
}

// customMethodNotAllowed implements http.Handler
func (c customMethodNotAllowed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	notImplementedAPIError(w)
}

// customNotFound implements http.Handler
func (c customNotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	notFoundAPIError(w)
}

func getRoutes() *httprouter.Router {
	// make a new httprouter
	router := httprouter.New()
	// declare all routes
	router.GET("/", rootHandler)
	// articles
	router.GET("/articles/", articlesHandler)
	router.GET("/articles/:id", articlesHandler)
	//categories
	router.GET("/categories/", categoriesHandler)
	router.GET("/categories/:id", categoriesHandler)

	router.NotFound = customNotFound{}
	router.MethodNotAllowed = customMethodNotAllowed{}

	// just make static public directory
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	return router
}

// Start func start server process and init all
// config server options
func Start() {
	// make just a new logger
	Logger = newLogger()
	// start the server
	err := http.ListenAndServe(domain+port, getRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
