package main_test

import (
	. "github.homedepot.com/joshq/db-broker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DbBroker", func() {
	It("breaks", func() {
		ThisIsPublic()
		Ω(1).
			Should(Equal(1))
	})
})
