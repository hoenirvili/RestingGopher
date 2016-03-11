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
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//
// rootHandler for root resource
//
func rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to my REST API build on GO !")

	defer func() {
		err := r.Body.Close()
		logIT(err)
	}()
}

//
// handler for articles resource
//
func articlesHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var (
		id uint64
	)

	// prepare header content-type
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// get id parsed and err if any
	id, err := resourceID(param.ByName("id"))

	// test the parsing scope
	switch err {
	// localhost:8080/articles/
	case errParamNotSet:
		// http method
		switch r.Method {
		case "GET":
			articlesGET(w)
		case "PUT":
			articlePUT(w, r)
		case "DELETE":
			articleDELETE(w, r)
		case "POST":
			articlePOST(w, r)
		default:
			internalAPIError(w)
		}
	// we have valid ID
	// localhost:8080/articles/{id}/
	case nil:
		switch r.Method {
		case "GET":
			articlesIDGET(w, id)
		case "PUT":
			articleIDPUT(w, r, id)
		case "DELETE":
			articleIDDELETE(w, r, id)
		case "POST":
		default:
			internalAPIError(w)
		}
	case errHighBitSet:
		//return response api to large number 64 int
		toLargeAPINumberError(w)
	default:
		// type assertion on strconv.NumError struct
		// if error is a type of this struct
		// check if attribute Err is strconv.ErrRange error
		numErr, ok := err.(*strconv.NumError)
		if ok {
			if numErr.Err == strconv.ErrRange {
				toLargeAPINumberError(w)
			}
		} else { // some other error that don't belongs
			// to strconv
			// internal service error api
			// parseINT error
			internalAPIError(w)
		}
	}

	defer func() {
		err := r.Body.Close()
		logIT(err)
	}()
}

//
//categoriesHandler for categories resource
//
func categoriesHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var (
		id uint64
	)

	// prepare header content-type
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// get id parsed and err if any
	id, err := resourceID(param.ByName("id"))

	// test the parsing scope
	switch err {
	// localhost:8080/categories/
	case errParamNotSet:
		// http method
		switch r.Method {
		case "GET":
			cateogoryGET(w)
		case "PUT":
			cateogoryPUT(w, r)
		case "DELETE":
			categoryDELETE(w, r)
		case "POST":
			categoryPOST(w, r)
		default:
			internalAPIError(w)
		}
	// we have valid ID
	// localhost:8080/categories/{id}/
	case nil:
		switch r.Method {
		case "GET":
			categoryIDGET(w, r, id)
		case "PUT":
			categoryIDPUT(w, r, id)
		case "DELETE":
			categoryIDDELETE(w, id)
		default:
			internalAPIError(w)
		}
	case errHighBitSet:
		//return response api to large number 64 int
		toLargeAPINumberError(w)
	// case strconv.ErrRange:
	// 	//return response api to large number 64 int
	// 	toLargeAPINumberError(w)
	default:
		// type assertion on strconv.NumError struct
		// if error is a type of this struct
		// check if attribute Err is strconv.ErrRange error
		numErr, ok := err.(*strconv.NumError)
		if ok {
			if numErr.Err == strconv.ErrRange {
				toLargeAPINumberError(w)
			}
		} else { // some other error that don't belongs
			// to strconv
			// internal service error api
			// parseINT error
			internalAPIError(w)
		}
	}

	defer func() {
		err := r.Body.Close()
		logIT(err)
	}()
}
