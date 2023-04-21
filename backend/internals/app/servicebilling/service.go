package servicebilling

import (
	"billing-service/db"
	"billing-service/internals/model"
	"billing-service/internals/repository"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	AddServiceBilling(p *ServicePayload) error
	AllServiceBillingDetails(time.Time, time.Time) ([]model.ServiceBilling, error)
}

type service struct {
	db *gorm.DB
}

func NewService() *service {
	return &service{db.DbManager()}
}

func (s *service) AddServiceBilling(p *ServicePayload) error {

	tx := s.db.Begin()

	payloadBytes, err := json.Marshal(p)
	if err != nil {
		return err
	}

	c := new(model.Customer)
	json.Unmarshal(payloadBytes, c)

	var customerId uint

	customerId = repository.GetCustomer(s.db, c.MobileNumber)

	if customerId == 0 {
		customerId, err = repository.AddCustomer(tx, c)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	serviceB := new(model.ServiceBilling)

	json.Unmarshal(payloadBytes, serviceB)

	serviceB.CustomerId = customerId

	err = repository.NewServiceBilling(tx, serviceB)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (s *service) AllServiceBillingDetails(startdate time.Time, enddate time.Time) ([]model.ServiceBilling, error) {

	data, err := repository.RetrieveServiceBillings(startdate, enddate, s.db)
	if err != nil {
		return nil, err
	}

	return data, nil
}
