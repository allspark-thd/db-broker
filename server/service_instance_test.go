package server_test

import (
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("/v2/service_instances/{guid}", func() {
	var server *httptest.Server

	BeforeEach(func() { server = createTestServer() })
	AfterEach(func() { server.Close() })
	// create => 202
})

// http://docs.cloudfoundry.org/services/api.html#polling
var _ = Describe(
	"GET /v2/service_instances/{guid}/last_operation",
	func() {

	})
