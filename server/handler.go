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

//
// rootHandler for root resource
//
func rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// w.Header().Add("Content-Type", "plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to my REST API build on GO !")
}

//
// handler for articles resource
//
func articlesHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	fmt.Fprintf(w, "artcile resource %s", param.ByName("id"))
}

//
//categoriesHandler for categories resource
//
func categoriesHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var (
		resourceQuery string
		id            uint64
	)
	// get id parsed and err if any
	id, err := resourceID(param.ByName("id"))

	// test the parsing scope
	switch err {
	case errNotSet:
		resourceQuery = "SELECT *FROM Category"
		// http method
		switch r.Method {
		case "GET":
			rows, err := Database.Query(resourceQuery)
			logIT(err)
			data, err := model.CategoriesJSON(rows)
			logIT(err)
			// prepare header content-type
			w.Header().Add("Content-Type", "application/json; charset=utf-8")

			// prepare  status code
			w.WriteHeader(http.StatusOK)

			// make newbody json response
			if _, err := w.Write(data); err != nil {
				//log server error
				Logger.Add("[GET] request on Categories\n Failed to write to response body\n [Query] " + resourceQuery)
			}
		default:
			internalAPIError(w)
		}

	case nil:
		resourceQuery = "SELECT *FROM Category WHERE ID_Category = ?"
		switch r.Method {
		case "GET":
			data, err := model.OneCategoryJSON(Database, resourceQuery, id)
			logIT(err)

			// prepare header content-type
			w.Header().Add("Content-Type", "application/json; charset=utf-8")

			// prepare  status code
			w.WriteHeader(http.StatusOK)

			// make newbody json response
			if _, err := w.Write(data); err != nil {
				//log server error
				Logger.Add("[GET] request on Categories\n Failed to write to response body\n [Query] " + resourceQuery)
			}
		default:
			internalAPIError(w)
		}

	case errHighBitSet:
		//return response api to large number 64 int
		toLargeAPINumberError(w)

	default:
		// internal service error api
		// parseINT error
		internalAPIError(w)
	}
}

//
// usersHandler for users resources
//
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
