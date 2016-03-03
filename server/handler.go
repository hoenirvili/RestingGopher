package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type customMethodNotAllowed struct {
}

// ServerHTTP implements http.Handler
func (c customMethodNotAllowed) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
func rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to my REST API build on GO !")
}

func articlesHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, param.ByName("id"))
	default:
		fmt.Fprintf(w, "Bad request")
	}
}

func categoriesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	switch r.Method {
	case "GET":
	default:
		internalAPIError(w)
	}
}
