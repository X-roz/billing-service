package repository

import (
	"billing-service/internals/model"

	"gorm.io/gorm"
)

func GetCustomer(db *gorm.DB, mobileNo string) uint {
	data := new(model.Customer)
	db.Model(&model.Customer{}).Where("mobile_number = ?", mobileNo).First(data)
	return data.ID
}

func AddCustomer(tx *gorm.DB, c *model.Customer) (uint, error) {
	err := tx.Create(c).Error
	return c.ID, err
}
