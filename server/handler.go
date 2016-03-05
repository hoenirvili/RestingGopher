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
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	// prepare header content-type
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// get id parsed and err if any
	id, err := resourceID(param.ByName("id"))

	// test the parsing scope
	switch err {
	case errParamNotSet:
		// http method
		switch r.Method {
		case "GET":
			resourceQuery = "SELECT *FROM Category"
			rows, err := Database.Query(resourceQuery)
			logIT(err)
			data, err := model.CategoriesJSON(rows)
			logIT(err)

			// prepare  status code
			w.WriteHeader(http.StatusOK)

			// make newbody json response
			if _, err := w.Write(data); err != nil {
				//log server error
				Logger.Add("[GET] request on Categories\n Failed to write to response body\n [Query] " + resourceQuery)
			}
		case "PUT":
			resourceQuery = "INSERT INTO Category VALUES(NULL, ?)"
			if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
				// new payload
				payload := CategoryPayload{}
				// decode into paylaod
				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
					Logger.Add("Can't decode json post category request")
				}
				// process the payload
				if payload.Data.ID != 0 && toHighSet(payload.Data.ID) {
					toLargeAPINumberError(w)
				} else { // make change post here
					// prepare query
					stmt, err := Database.Prepare(resourceQuery)
					logIT(err) // log it if err
					// exec stmt after prepare
					err = Database.ExecStmt(stmt, payload.Data.Name)
					logIT(err) // log it if err
					// select the new row created
					rows, err := Database.Query("Select *FROM Category WHERE Name = ? ", payload.Data.Name)
					logIT(err) // log it if err
					// if one/many rows are the same return it
					data, err := model.CategoriesJSON(rows)
					logIT(err) // log it if err
					// prepare status codes
					w.WriteHeader(http.StatusCreated)
					// write json response
					if _, err := w.Write(data); err != nil {
						//log server error
						Logger.Add("[PUT] request on Categories\n Failed to write to response body\n [Query] " + resourceQuery)
					}
				}
				defer func() {
					err := r.Body.Close()
					logIT(err)
				}()
			} else { // bad content type request
				appropriateHeaderError(w)
			}
			// case ""
		default:
			internalAPIError(w)
		}
	// we have valid ID
	case nil:
		switch r.Method {
		case "GET":
			resourceQuery = "SELECT *FROM Category WHERE ID_Category = ?"
			data, err := model.OneCategoryJSON(Database, resourceQuery, id)
			//think that we have a good response back from sql
			status := http.StatusOK
			// if the response if empty
			if err == model.ErrNoContent {
				// prepare  status code
				status = http.StatusNotFound
			} else { // other error
				logIT(err)
			}
			// prepare  status code
			w.WriteHeader(status)

			// make newbody json response
			if _, err := w.Write(data); err != nil {
				//log server error
				Logger.Add("[GET] request on Categories\n Failed to write to response body\n [Query] " + resourceQuery)
			}

		case "PUT":
			resourceQuery = "UPDATE Category SET  Name = ? WHERE ID_Category = ?"
			if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
				// new payload
				payload := CategoryPayload{}
				// decode into paylaod
				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
					Logger.Add("Can't decode json post category request")
				}
				// process the payload
				if payload.Data.ID != 0 && toHighSet(payload.Data.ID) {
					toLargeAPINumberError(w)
				} else { // make change post here
					// prepare query
					stmt, err := Database.Prepare(resourceQuery)
					logIT(err) // log it if err
					// exec stmt after prepare
					err = Database.ExecStmt(stmt, payload.Data.Name, id)
					logIT(err) // log it if err
					// select the new row created
					rows, err := Database.Query("Select *FROM Category WHERE Name = ? ", payload.Data.Name)
					logIT(err) // log it if err
					// if one/many rows are the same return it
					data, err := model.CategoriesJSON(rows)
					logIT(err) // log it if err
					// prepare status codes
					w.WriteHeader(http.StatusCreated)
					// write json response
					if _, err := w.Write(data); err != nil {
						//log server error
						Logger.Add("[PUT] request on Categories\n Failed to write to response body\n [Query] " + resourceQuery)
					}
				}
				defer func() {
					err := r.Body.Close()
					logIT(err)
				}()
			} else { // bad content type request
				appropriateHeaderError(w)
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
