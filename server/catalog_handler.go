package server

import (
	"net/http"

	"github.homedepot.com/joshq/db-broker/model"
)

// NewCatalogHandler creates a handler that
// outputs the json list of services provided
// by the service broker
func NewCatalogHandler() http.Handler {
	return catalogHandler{}
}

type catalogHandler struct{}

func (catalogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeJSON(w,
		model.Catalog{
			Services: []model.Service{
				model.Service{},
			},
		})
}
