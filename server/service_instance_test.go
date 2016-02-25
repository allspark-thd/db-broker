package server_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Put(urlStr string, body string) (*http.Response, error) {
	req, _ := http.NewRequest("PUT", urlStr, bytes.NewBufferString(body))
	return http.DefaultClient.Do(req)
}

var _ = Describe("/v2/service_instances/{guid}", func() {
	var server *httptest.Server
	BeforeEach(func() {
		server = createTestServer()
	})
	AfterEach(func() { server.Close() })

	var instanceURL = func(ID string) string {
		return fmt.Sprintf("%s/v2/service_instances/%s", server.URL, ID)
	}
	const svcjson = `
	{
		"organization_guid": "org-guid-here",
		"plan_id":           "plan-guid-here",
		"service_id":        "service-guid-here",
		"space_guid":        "space-guid-here"
	}`
	Describe("PUT", func() {
		It("requires `?accepts_incomplete=true`", func() {
			response, _ := Put(instanceURL("org-guid-here"), svcjson)

			Ω(response.StatusCode).
				Should(Equal(422), "Unprocessable Entity")

			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			Ω(body).Should(
				MatchJSON(`{
					"error": "AsyncRequired",
					"description": "This service plan requires client support for asynchronous service operations."
					}`))
		})
		It("saves a service instance", func() {
			url := fmt.Sprintf(
				"%s?accepts_incomplete=true",
				instanceURL("org-guid-here"))
			response, _ := Put(url, svcjson)
			Ω(response.StatusCode).Should(Equal(http.StatusAccepted))
		})
	})
	Describe("GET", func() {
		BeforeEach(func() {
			url := fmt.Sprintf(
				"%s?accepts_incomplete=true",
				instanceURL("org-guid-here"))
			Put(url, svcjson)
		})

		It("looks up service instance by guid", func() {
			response, _ := http.Get(instanceURL("org-guid-here"))
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			Ω(body).Should(MatchJSON(svcjson))
		})
	})
	// create => 202
})

// http://docs.cloudfoundry.org/services/api.html#polling
var _ = Describe(
	"GET /v2/service_instances/{guid}/last_operation",
	func() {

	})
