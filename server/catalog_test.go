package server_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.homedepot.com/joshq/db-broker/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("/v2/catalog", func() {
	var server *httptest.Server
	var response *http.Response

	BeforeEach(func() {
		server = createTestServer()
		response, _ = http.Get(server.URL + "/v2/catalog")
	})
	AfterEach(func() { server.Close() })

	It("responds to GET", func() {
		Ω(response.StatusCode).Should(Equal(200))
	})

	It("has content-type `application/json`", func() {
		Ω(response.Header.Get("content-type")).
			Should(ContainSubstring("application/json"))
	})

	It("returns a catalog", func() {
		catalog := parseCatalog(response)
		Ω(catalog).ShouldNot(BeNil())
		Ω(catalog.Services).ShouldNot(BeEmpty())
	})

	// TODO
	It("has valid service definitions", func() {})
})

func parseCatalog(response *http.Response) *Catalog {
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	catalog := &Catalog{}
	json.Unmarshal(body, catalog)
	return catalog
}
