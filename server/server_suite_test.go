package server_test

import (
	"net/http/httptest"

	"testing"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.homedepot.com/joshq/db-broker/dao"
	. "github.homedepot.com/joshq/db-broker/server"

	_ "github.com/mattn/go-sqlite3"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

func createTestServer() *httptest.Server {
	db, _ := gorm.Open("sqlite3", ":memory:")
	dao := dao.NewBrokerDAO(db)
	return httptest.NewServer(NewRouter(dao))
}

// type closer func()

func close(server *httptest.Server) func() {
	return func() {
		server.Close()
	}
}
