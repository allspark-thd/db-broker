package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter returns an http.Handler
// that implements the cloud foundry
// service broker contract
func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.Handle("/v2/catalog", NewCatalogHandler())
	return router
}
