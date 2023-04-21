package repository

import (
	"billing-service/internals/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewSalesBilling(tx *gorm.DB, s *model.SalesBilling) error {
	return tx.Create(s).Error
}

func RetrieveSalesBillings(startdate time.Time, enddate time.Time, db *gorm.DB) ([]model.SalesBilling, error) {
	data := new([]model.SalesBilling)
	err := db.Model(&model.SalesBilling{}).Where("created_at between ? and ?", startdate, enddate).Preload("SalesDetails.Product").Preload(clause.Associations).Find(data).Error
	return *data, err
}
