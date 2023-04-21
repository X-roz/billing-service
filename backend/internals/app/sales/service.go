package sales

import (
	"billing-service/db"
	"billing-service/internals/model"
	"billing-service/internals/repository"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	AddSalesBilling(*SalesPayload) error
	AllSalesBillingDetails(time.Time, time.Time) ([]model.SalesBilling, error)
}

type service struct {
	db *gorm.DB
}

func NewService() *service {
	return &service{db.DbManager()}
}

func (s *service) AddSalesBilling(p *SalesPayload) error {

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

	sales := new(model.SalesBilling)

	json.Unmarshal(payloadBytes, sales)

	sales.CustomerId = customerId

	err = repository.NewSalesBilling(tx, sales)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, d := range p.SalesDetails {
		err := repository.UpdateQtyBasedOnBilling(tx, d.ProductId, d.OrderedQuantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (s *service) AllSalesBillingDetails(startdate time.Time, enddate time.Time) ([]model.SalesBilling, error) {

	data, err := repository.RetrieveSalesBillings(startdate, enddate, s.db)
	if err != nil {
		return nil, err
	}

	return data, nil
}
