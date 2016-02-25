package dao_test

import (
	"os"

	. "github.homedepot.com/joshq/db-broker/dao"
	. "github.homedepot.com/joshq/db-broker/model"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jinzhu/gorm"
)

var _ = Describe("BrokerDAO", func() {
	Describe("service instances", testServiceInstances)
})

func testServiceInstances() {
	var dao BrokerDAO
	var testInstance ServiceInstance
	BeforeEach(func() {
		db, _ := gorm.Open("sqlite3", "broker.sqlite")
		dao = NewBrokerDAO(db)
		testInstance = ServiceInstance{
			ID:               "svc-instance-id",
			OrganizationGuid: "org-guid",
			PlanId:           "svc-plan-id",
			ServiceId:        "svc-guid",
		}
	})
	AfterEach(func() {
		os.Remove("broker.sqlite")
	})
	Describe("#SaveServiceInstance", func() {
		It("saves new instances", func() {
			Ω(dao.SaveServiceInstance(testInstance)).
				ShouldNot(HaveOccurred())
		})
		It("overwrites existing", func() {
			// add test instance
			dao.SaveServiceInstance(testInstance)
			before, _ := dao.FindServiceInstance("svc-instance-id")

			testInstance.OrganizationGuid = "new-org-id"
			dao.SaveServiceInstance(testInstance)

			after, _ := dao.FindServiceInstance("svc-instance-id")

			Ω(before.OrganizationGuid).
				Should(Equal("org-guid"))
			Ω(after.OrganizationGuid).
				Should(Equal("new-org-id"))
		})
	})
	Describe("#FindServiceInstance", func() {
		It("returns nil for non-existent", func() {
			instance, err := dao.FindServiceInstance("does-not-exist")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(instance).Should(BeNil())
		})

		It("finds by id", func() {
			dao.SaveServiceInstance(testInstance)
			instance, err := dao.FindServiceInstance("svc-instance-id")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(instance).ShouldNot(BeNil())
			Ω(instance.ID).
				Should(Equal("svc-instance-id"),
				"found correct instance")
		})
	})
}
