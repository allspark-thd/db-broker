package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.homedepot.com/joshq/db-broker/dao"
	"github.homedepot.com/joshq/db-broker/model"
)

/*
router *mux.Router
func AddRoutingHandler(router *mux.Router, api gmt.GmtAPI) {
s.Methods("GET").HandlerFunc(handler.getPrimaryCell)
s.Methods("PUT").HandlerFunc(handler.setPrimaryCell)
*/

// NewServiceInstanceHandler is the default handler for service instance
// management routes.
func NewServiceInstanceHandler(router *mux.Router, brokerDao dao.BrokerDAO) {
	handler := serviceInstanceHandler{
		BrokerDAO: brokerDao,
	}

	r := router.Path("/{guid}").Subrouter()
	r.Methods("GET").Handler(
		http.HandlerFunc(handler.GET))
	r.Methods("PUT").Handler(
		http.HandlerFunc(handler.PUT))
}

type serviceInstanceHandler struct {
	dao.BrokerDAO
}

func (h serviceInstanceHandler) GET(w http.ResponseWriter, r *http.Request) {
	guid := mux.Vars(r)["guid"]
	log.Printf("Fetching service instance `%s`", guid)
	svcInstance, _ := h.FindServiceInstance(guid)
	writeJSON(w, svcInstance)
}

func (h serviceInstanceHandler) PUT(w http.ResponseWriter, r *http.Request) {
	var instance model.ServiceInstance
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &instance)

	instance.ID = mux.Vars(r)["guid"]

	if r.URL.Query().Get("accepts_incomplete") != "true" {
		w.WriteHeader(422)
		writeJSON(w, map[string]string{
			"error": "AsyncRequired",
			"description": strings.Join([]string{
				`This service plan requires client`,
				`support for asynchronous service operations.`}, " "),
		})
		return
	}

	log.Printf("Saving service instance `%s`", instance.ID)
	h.BrokerDAO.SaveServiceInstance(instance)
	w.WriteHeader(http.StatusAccepted)
	writeJSON(w, map[string]string{})
}
