package dao

import (
	"github.com/jinzhu/gorm"
	"github.homedepot.com/joshq/db-broker/model"
)

//BrokerDAO is a contract for interacting with
//the service broker's underlying data store
//for service instances and bindings.
type BrokerDAO interface {
	SaveServiceInstance(instance model.ServiceInstance) error
	FindServiceInstance(svcID string) (*model.ServiceInstance, error)
}

type daoimpl struct {
	gorm.DB
	model.Catalog
}

//NewBrokerDAO is the default implementation of the
//BrokerDAO interface.
func NewBrokerDAO(db gorm.DB) BrokerDAO {
	db.AutoMigrate(&model.ServiceInstance{})
	return daoimpl{
		DB: db,
	}
}

func (dao daoimpl) SaveServiceInstance(svc model.ServiceInstance) error {
	dao.Create(svc)
	return dao.Save(svc).Error
}

func (dao daoimpl) FindServiceInstance(svcID string) (
	*model.ServiceInstance, error) {
	var instance model.ServiceInstance
	dao.DB.Find(&instance, model.ServiceInstance{ID: svcID})

	if instance.ID != svcID {
		// not found
		return nil, dao.DB.Error
	}
	return &instance, dao.DB.Error
}
