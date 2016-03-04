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
	"fmt"
	"net/http"

	"github.com/hoenirvili/RestingGopher/model"
	"github.com/julienschmidt/httprouter"
)

// handler for root resource
func rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// w.Header().Add("Content-Type", "plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to my REST API build on GO !")
}

// handler for articles resource
func articlesHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var id string

	id = param.ByName("id")

	switch r.Method {
	case "GET":
		// retrive info about one article
		if len(id) > 0 {
			fmt.Fprintf(w, id)
		} else { // etrive all articles in a simple page order.

		}
	case "POST":

	case "PUT":
	case "DELETE":
	default:
		// any other http verbs
		notImplementedAPIError(w)
	}
}

//handler for Categories resource
func categoriesHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var id string
	id = param.ByName("id")
	switch r.Method {
	case "GET":
		// retrive specific category info
		if len(id) > 0 {
			categoryJSON := model.Categories{}
			err := Db.Query("Select *from Categories;", categoryJSON)
			if err != nil {
				Logger.Add(err.Error())
			}

		} else { // retrive all categories
		}
	case "POST":
	case "PUT":
	case "DELETE":
	default:
		// any other http verbs
		notImplementedAPIError(w)
	}
}

// handler for Users resources
func usersHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	switch r.Method {
	case "GET":
	case "POST":
	case "PUT":
	case "DELETE":
	default:
		notImplementedAPIError(w)
	}
}
