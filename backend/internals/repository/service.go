package repository

import (
	"billing-service/internals/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewServiceBilling(tx *gorm.DB, s *model.ServiceBilling) error {
	return tx.Create(s).Error
}

func RetrieveServiceBillings(startdate time.Time, enddate time.Time, db *gorm.DB) ([]model.ServiceBilling, error) {
	data := new([]model.ServiceBilling)
	err := db.Model(&model.ServiceBilling{}).Where("created_at between ? and ?", startdate, enddate).Preload(clause.Associations).Find(data).Error
	return *data, err
}
