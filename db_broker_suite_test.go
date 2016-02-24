package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDbBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DbBroker Suite")
}
