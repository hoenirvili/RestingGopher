package server

import (
	"encoding/json"
	"net/http"
)

// DefaultInternalError default byte type
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

// ErrServerJSON playload does not implment the error interface
// because we use just to fill the http response body
type ErrServerJSON struct {
	Data ErrData
}

// NewErrServerJSON returns a new ErrServerJSON struct
// for marshaling json
func NewErrServerJSON(msg string) ErrServerJSON {
	return ErrServerJSON{
		Data: ErrData{Message: msg},
	}
}

// JSON func returns indent json serialized
// It's a wrapper around MarshalIndent
// using just "" and " " for indentation
func (e ErrServerJSON) JSON() []byte {
	body, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		return []byte(DefaultInternalError)
	}
	return body
}

// internalApiError sends JSON response with error message and header 500
func internalAPIError(w http.ResponseWriter) {
	// prepare header content type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// prepare  header status code
	w.WriteHeader(http.StatusInternalServerError) // 500
	// Make new body json response
	body := NewErrServerJSON("Internal Server Error test").JSON()
	if _, err := w.Write(body); err != nil {
		Logger.Add("Internal API Error Write to response body failed")
	}
}
