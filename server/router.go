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
	"net/http"

	"github.com/julienschmidt/httprouter"
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

// getRoutes func returns a new custom httprouter
func getRoutes() *httprouter.Router {
	// make a new httprouter
	router := httprouter.New()

	// root router part
	router.GET("/", rootHandler)
	// articles router parts
	router.GET("/articles/", articlesHandler)
	router.GET("/articles/:id", articlesHandler)
	router.PUT("/articles/", articlesHandler)
	router.PUT("/articles/:id", articlesHandler)

	//TODO
	router.DELETE("/articles/", articlesHandler)
	router.DELETE("/articles/:id", articlesHandler)

	// categories router parts
	router.GET("/categories/", categoriesHandler)
	router.GET("/categories/:id", categoriesHandler)
	router.PUT("/categories/", categoriesHandler)
	router.PUT("/categories/:id", categoriesHandler)
	router.DELETE("/categories/", categoriesHandler)
	router.DELETE("/categories/:id", categoriesHandler)
	router.POST("/categories/", categoriesHandler)

	// init custom notfound/notallowed methods
	router.NotFound = customNotFound{}
	router.MethodNotAllowed = customMethodNotAllowed{}

	// just make static public directory
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	return router
}
