package server_test

import (
	"net/http/httptest"

	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.homedepot.com/joshq/db-broker/server"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

func createTestServer() *httptest.Server {
	return httptest.NewServer(NewRouter())
}

// type closer func()

func close(server *httptest.Server) func() {
	return func() {
		server.Close()
	}
}
