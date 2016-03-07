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
	const resourceQuery = "SELECT *FROM Category"
	rows, err := Database.Query(resourceQuery)
	logIT(err)
	data, err := model.CategoriesJSON(rows)
	logIT(err)

	// prepare  status code
	// w.WriteHeader(http.StatusOK)

	// make newbody json response
	if _, err := w.Write(data); err != nil {
		//log server error
		Logger.Add("[GET] request on Categories\n Failed to write to response body\n [Query] " + resourceQuery)
	}
}

func categoryIDGET(w http.ResponseWriter, r *http.Request, id uint64) {
	const resourceQuery = "SELECT *FROM Category WHERE ID_Category = ?"
	data, err := model.OneCategoryJSON(Database, resourceQuery, id)
	// if the response if empty
	if err == model.ErrNoContent {
		notFoundAPIError(w)
	} else { // other error
		logIT(err)
	}
	// prepare  status code
	// w.WriteHeader(http.StatusOK)

	// make newbody json response
	if _, err := w.Write(data); err != nil {
		//log server error
		Logger.Add("[GET] request on Categories ID\n Failed to write to response body\n [Query] " + resourceQuery)
	}

}
func cateogoryPUT(w http.ResponseWriter, r *http.Request) {
	const resourceQuery = "INSERT INTO Category VALUES(NULL, ?)"
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		// new payload
		payload := CategoryPayload{}
		// decode into paylaod
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || payload.Data.ID == 0 || payload.Data.Name == "" {
			Logger.Add("Can't decode json put category request")
			invalidJSONFormatError(w)
			goto end
		}
		// process the payload
		if toHighSet(payload.Data.ID) {
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
end:
}
func categoryIDPUT(w http.ResponseWriter, r *http.Request, id uint64) {
	const resourceQuery = "UPDATE Category SET Name = ? WHERE ID_Category = ?"
	// if the content we want to PUT is in JSON format that means we can
	// first process it
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		// create a new CategoryPayload to hold the data from
		// http.Request
		payload := CategoryPayload{}
		// we must deode the http.Request JSON into the payload struct
		err := json.NewDecoder(r.Body).Decode(&payload)
		// if the payload can't be decoded
		// sent user api error , log it and jump to the end of the func
		if err != nil || payload.Data.Name == "" {
			Logger.Add("Can't PUT id json category request")
			invalidJSONFormatError(w)
			goto end
		}
		// if the decoding process was complete
		// we need to check one more thing that the high
		// id bit is not set that 64 bit
		if toHighSet(payload.Data.ID) { // if it's set just return error to the user
			toLargeAPINumberError(w)
		} else { // else the bit is not set we must insert/update id field into database

			// Select the location were we want to update/insert our data
			_, err := model.OneCategoryJSON(Database, "SELECT *FROM Category WHERE ID_Category = ?", id)
			// if we found the content
			// that means we have a row returned
			// we must update
			if err != model.ErrNoContent {
				// UPDATE STMT
				stmt, err := Database.Prepare(resourceQuery)
				logIT(err) // log it if err
				// exec stmt after prepare
				err = Database.ExecStmt(stmt, payload.Data.Name, id)
				logIT(err) // log it if err
			} else { // we didn't found some content just insert into the missing content
				stmt, err := Database.Prepare("INSERT INTO Category VALUES( ? , ? )")
				logIT(err)
				err = Database.ExecStmt(stmt, id, payload.Data.Name)
				logIT(err)
			}
			// after we done updating/inserting the new content
			// we must select the content created and return it into
			// a json response  with http status 201 if we inserted(created), or
			// altered with 202
			data, err := model.OneCategoryJSON(Database, "Select *FROM Category WHERE ID_Category = ? ", id)
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
end:
}

func categoryDELETE(w http.ResponseWriter, r *http.Request) {
	const resourceQuery = "DELETE FROM Category WHERE ID_Category = ? && Name = ? "
	// if request header has json
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		// new payload
		payload := CategoryPayload{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || payload.Data.ID == 0 || payload.Data.Name == "" {
			Logger.Add("Can't decode json delete category request")
			invalidJSONFormatError(w)
			goto end
		}

		if toHighSet(payload.Data.ID) {
			toLargeAPINumberError(w)
		} else { // make change post here
			_, err := model.OneCategoryJSON(Database, "SELECT *from Category WHERE ID_Category = ? && Name = ?", payload.Data.ID, payload.Data.Name)
			if err == model.ErrNoContent {
				notFoundAPIError(w)
				goto end
			}
			// if we found the entry we want to delete
			// prepare stmt for delete
			stmt, err := Database.Prepare(resourceQuery)
			logIT(err)
			// exec DELETE Stmt
			err = Database.ExecStmt(stmt, payload.Data.ID, payload.Data.Name)
			logIT(err)

			// prepare status codes
			w.WriteHeader(http.StatusOK) // 200
			//make anon struct for response JSON
			responseJSON := struct {
				Data struct {
					ID      uint64 `json: "ID"`
					Name    string `json: "Name"`
					Message string `json: "Message"`
				}
			}{}
			responseJSON.Data.ID = payload.Data.ID
			responseJSON.Data.Name = payload.Data.Name
			responseJSON.Data.Message = "DELETE request successful"

			// marshall anon struct
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

func categoryIDDELETE(w http.ResponseWriter, id uint64) {
	const resourceQuery = "DELETE FROM Category WHERE ID_Category = ?"
	if toHighSet(id) {
		toLargeAPINumberError(w)
	} else {
		stmt, err := Database.Prepare(resourceQuery)
		logIT(err)
		err = Database.ExecStmt(stmt, id)
		logIT(err)
	}

	responseJSON := struct {
		Data struct {
			ID      uint64 `json: "ID_Resource"`
			Message string `json: "Message"`
		}
	}{}
	responseJSON.Data.ID = id
	responseJSON.Data.Message = "Successful delete resource"
	data, err := json.MarshalIndent(responseJSON, "", " ")
	logIT(err)

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		//log server error
		Logger.Add("[DELETE] request on Categories \n Failed to write to response body\n [Query] " + resourceQuery)
	}
}
func categoryPOST(w http.ResponseWriter, r *http.Request) {
	const (
		resourceQuery = "INSERT INTO Category VALUES( ?, ? )"
		updateQuery   = "UPDATE Category SET Name = ? WHERE ID_Category = ? || Name = ?"
	)
	var message string
	// we have JSON POST request
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		// new payload
		payload := CategoryPayload{}
		// decode into payload struct
		err := json.NewDecoder(r.Body).Decode(&payload)
		// if the payload is not ok deoced
		if err != nil && payload.Data.ID == 0 || payload.Data.Name == "" {
			invalidJSONFormatError(w)
			Logger.Add("Can't decode json respons POST category request")
			goto end
		}
		// process ID payload if it's ok
		if toHighSet(payload.Data.ID) {
			toLargeAPINumberError(w)
		} else { // process the request
			_, err := model.OneCategoryJSON(Database, "SELECT *from Category WHERE ID_Category= ? || Name = ?", payload.Data.ID, payload.Data.Name)
			// if we found the content already
			// that means we have a row returned
			// we must update it not insert duplicate
			if err != model.ErrNoContent {
				// UPDATE STMT
				err := Database.Exec(updateQuery, payload.Data.Name, payload.Data.ID, payload.Data.Name)
				logIT(err)
				message = "Successful modify resource "
			} else { // else we don't have the data and we must inserted
				stmt, err := Database.Prepare(resourceQuery)
				logIT(err)
				err = Database.ExecStmt(stmt, payload.Data.ID, payload.Data.Name)
				logIT(err)
				message = "Successful inserted resource "
			}
			w.WriteHeader(http.StatusCreated)
			responseJSON := struct {
				Data struct {
					ID      uint64 `json: "ID_Resource"`
					Name    string `json: "Name"`
					Message string `json: "Message"`
				}
			}{}
			responseJSON.Data.ID = payload.Data.ID
			responseJSON.Data.Name = payload.Data.Name
			responseJSON.Data.Message = message
			if err := json.NewEncoder(w).Encode(responseJSON); err != nil {
				Logger.Add("[POST] request on Cateogires \n Failed to write response body\n [Query] " + resourceQuery)
			}
		}
	}
end:
}

func articlesGET(w http.ResponseWriter) {

}

func articlesIDGET(w http.ResponseWriter, id uint64) {

}
