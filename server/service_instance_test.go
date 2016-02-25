package server_test

import (
	"bytes"
	"encoding/json"
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
		"organization_guid": "instnc-guid-here",
		"plan_id":           "plan-guid-here",
		"service_id":        "service-guid-here",
		"space_guid":        "space-guid-here"
	}`
	Describe("PUT", func() {
		It("requires `?accepts_incomplete=true`", func() {
			response, _ := Put(instanceURL("instnc-guid-here"), svcjson)

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
				instanceURL("instnc-guid-here"))
			response, _ := Put(url, svcjson)
			Ω(response.StatusCode).Should(Equal(http.StatusAccepted))
		})
		// TODO test that we contact provisioner api
	})
	Describe("GET", func() {
		BeforeEach(func() {
			url := fmt.Sprintf(
				"%s?accepts_incomplete=true",
				instanceURL("instnc-guid-here"))
			Put(url, svcjson)
		})

		It("404's for non-existent instance", func() {
			response, _ := http.Get(instanceURL("non-existent-guid"))
			Ω(response.StatusCode).
				Should(Equal(http.StatusNotFound))
		})

		It("looks up service instance by guid", func() {
			response, _ := http.Get(instanceURL("instnc-guid-here"))
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			Ω(body).Should(MatchJSON(svcjson))
		})

		Describe("/last_operation", func() {
			var getLastOp = func() map[string]string {
				response, _ := http.Get(instanceURL("instnc-guid-here"))
				defer response.Body.Close()
				body, _ := ioutil.ReadAll(response.Body)
				var statemap map[string]string
				json.Unmarshal(body, &statemap)
				return statemap
			}

			// TODO
			It("checks credentials with the Todd", func() {
				getLastOp()
			})
		})
	})
})

// http://docs.cloudfoundry.org/services/api.html#polling
