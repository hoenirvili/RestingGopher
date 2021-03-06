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
	"errors"
	"net/http"
	"strings"
)

var (
	errHighBitSet  = errors.New("High bit set")
	errParamNotSet = errors.New("Params not set")
)

// DefaultInternalError JSON string response
const (
	DefaultInternalError = `{
  "Data:" {
   "Message":"The server has encountered some internal difficulty, like some bad sql connection pooling timeout, bad request,error login, please try again later"
  }
}`
)

// ErrData hold the body json
// that will be written to the ReponseWriter
type ErrData struct {
	Message string
}

// ErrServer playload does not implment the error interface
// because we use just to fill the http response body
type ErrServer struct {
	Data ErrData
}

// NewErrServer returns a new ErrServer struct
// for marshaling json
func NewErrServer(msg string) ErrServer {
	return ErrServer{
		Data: ErrData{Message: msg},
	}
}

// JSON func returns indent json serialized
// It's a wrapper around MarshalIndent
// using just "" and " " for indentation
func (e ErrServer) JSON() []byte {
	body, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		return []byte(DefaultInternalError)
	}
	return body
}

// internalApiError sends JSON response with error message and header 500
func internalAPIError(w http.ResponseWriter) {
	// is header is not set
	if !(strings.Contains(w.Header().Get("Content-Type"),
		"application/json")) {
		// prepare header content type
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	// prepare   status code
	w.WriteHeader(http.StatusInternalServerError) // 500
	// Make new body json response
	body := NewErrServer("The server has encountered an error try again later.").JSON()
	if _, err := w.Write(body); err != nil {
		Logger.Add("Internal API Error Write to response body failed")
	}
}

// notFoundResource
func notFoundAPIError(w http.ResponseWriter) {
	// if header is not set
	if !(strings.Contains(w.Header().Get("Content-Type"), "application/json")) {
		// prepare header content-type
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
	}
	// prepare  status code
	w.WriteHeader(http.StatusNotFound) //404
	// Make newbody json response
	body := NewErrServer("Content not found, the requested resource for reading,writing,modifing can't be accessed").JSON()
	if _, err := w.Write(body); err != nil {
		Logger.Add("Content not found Write to response body failed")
	}
}

func notImplementedAPIError(w http.ResponseWriter) {
	// if header is not set
	if !(strings.Contains(w.Header().Get("Content-Type"), "application/json")) {
		// prepare header content-type
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
	}
	// prepare  status code
	w.WriteHeader(http.StatusNotImplemented) //503
	// make newbody json response
	body := NewErrServer("This resource or content is not implemented yet").JSON()
	if _, err := w.Write(body); err != nil {
		Logger.Add("Resource or content not implemented yet Write to response body failed")
	}
}

func toLargeAPINumberError(w http.ResponseWriter) {
	// if header is not set
	if !(strings.Contains(w.Header().Get("Content-Type"), "application/json")) {
		// prepare header content-type
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
	}
	// prepare  status code
	w.WriteHeader(http.StatusNotAcceptable) //406
	// make newbody json response
	body := NewErrServer("The number of the resource passed is to large.API can't hadle numbers larger than 9223372036854775807").JSON()
	if _, err := w.Write(body); err != nil {
		Logger.Add("Number of the resource passed is to large Write to response body failed")
	}
}

func appropriateHeaderError(w http.ResponseWriter) {
	// if header is not set
	if !(strings.Contains(w.Header().Get("Content-Type"), "application/json")) {
		// prepare header content-type
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
	}
	// prepare  status code
	w.WriteHeader(http.StatusNotAcceptable) //406
	// make newbody json response
	body := NewErrServer("Please enter a valid Content-Type ! This api just implements JSON request/response").JSON()
	if _, err := w.Write(body); err != nil {
		Logger.Add("Can't write to response , content-type api error")
	}
}

func invalidJSONFormatError(w http.ResponseWriter) {
	// if header is not set
	if !(strings.Contains(w.Header().Get("Content-Type"), "application/json")) {
		// prepare header content-type
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
	}
	// prepare  status code
	w.WriteHeader(http.StatusBadRequest)
	// make newbody json response
	body := NewErrServer("Invalid JSON format error, please consult the documentation of this rest api").JSON()
	if _, err := w.Write(body); err != nil {
		Logger.Add("Can't write to response , content-type api error")
	}
}
