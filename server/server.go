package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.homedepot.com/joshq/db-broker/dao"
)

// NewRouter returns an http.Handler
// that implements the cloud foundry
// service broker contract
func NewRouter(brokerDAO dao.BrokerDAO) http.Handler {
	router := mux.NewRouter()
	router.Handle("/v2/catalog", NewCatalogHandler())

	NewServiceInstanceHandler(
		router.PathPrefix("/v2/service_instances").Subrouter(),
		brokerDAO)

	return router
}

//router.PathPrefix(path).Subrouter()
