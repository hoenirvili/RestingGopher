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
	"net/http"
	"strings"

	"github.com/hoenirvili/RestingGopher/model"
)

func cateogoryGET(w http.ResponseWriter) {
	resourceQuery := "SELECT *FROM Category"
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
}

func categoryIDGET(w http.ResponseWriter, r *http.Request, id uint64) {
	resourceQuery := "SELECT *FROM Category WHERE ID_Category = ?"
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
		Logger.Add("[GET] request on Categories ID\n Failed to write to response body\n [Query] " + resourceQuery)
	}

}
func cateogoryPUT(w http.ResponseWriter, r *http.Request) {
	resourceQuery := "INSERT INTO Category VALUES(NULL, ?)"
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		// new payload
		payload := CategoryPayload{}
		// decode into paylaod
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			Logger.Add("Can't decode json put category request")
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

	} else { // bad content type request
		appropriateHeaderError(w)
	}

}
func categoryIDPUT(w http.ResponseWriter, r *http.Request, id uint64) {
	resourceQuery := "UPDATE Category SET  Name = ? WHERE ID_Category = ?"
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		// new payload
		payload := CategoryPayload{}
		// decode into paylaod
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			Logger.Add("Can't decode json put category request")
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
				Logger.Add("[PUT] request on Categories ID\n Failed to write to response body\n [Query] " + resourceQuery)
			}
		}
	} else { // bad content type request
		appropriateHeaderError(w)
	}

}

func categoryDELETE(w http.ResponseWriter, r *http.Request) {
	//TODO FIX DELETE
	resourceQuery := "DELETE FROM Category WHERE ID_Category = ? && Name = ? "
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		// new payload
		payload := CategoryPayload{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			Logger.Add("Can't decode json delete category request")
			invalidJSONFormatError(w)
			goto end
		}

		if payload.Data.ID != 0 && toHighSet(payload.Data.ID) {
			toLargeAPINumberError(w)
		} else { // make change post here
			stmt, err := Database.Prepare(resourceQuery)
			logIT(err)
			err = Database.ExecStmt(stmt, payload.Data.Name, payload.Data.ID)
			logIT(err)
			// prepare status codes
			w.WriteHeader(http.StatusOK)
			// write json response

			responseJSON := struct {
				Data struct {
					ID      uint64
					Name    string
					Message string
				}
			}{}
			responseJSON.Data.ID = payload.Data.ID
			responseJSON.Data.Name = payload.Data.Name
			responseJSON.Data.Message = "DELETE request successful"

			data, err := json.MarshalIndent(responseJSON, "", " ")
			logIT(err)

			if _, err := w.Write(data); err != nil {
				//log server error
				Logger.Add("[DELETE] request on Categories \n Failed to write to response body\n [Query] " + resourceQuery)
			}
		}
	} else { // bad content type request
		appropriateHeaderError(w)
	}

end:
}
