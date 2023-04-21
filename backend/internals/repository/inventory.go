package repository

import (
	"billing-service/db"
	"billing-service/internals/model"

	"gorm.io/gorm"
)

type ItemsRepository interface {
	AddItems(*model.Inventory) error
	GetItems(interface{}) (*[]model.Inventory, error)
	UpdateItem(interface{}, *model.Inventory) error
	UpdateQty(interface{}, int64) error
}

type itemsConn struct {
	db *gorm.DB
}

func NewItemsConn() *itemsConn {
	return &itemsConn{db.DbManager()}
}

func (conn *itemsConn) AddItems(data *model.Inventory) error {
	return conn.db.Model(&model.Inventory{}).Create(data).Error
}

func (conn *itemsConn) GetItems(query interface{}) (*[]model.Inventory, error) {
	respData := new([]model.Inventory)
	err := conn.db.Model(&model.Inventory{}).Where(query).Find(respData).Error
	return respData, err
}

func (conn *itemsConn) UpdateItem(query interface{}, data *model.Inventory) error {
	return conn.db.Model(&model.Inventory{}).Where(query).Updates(data).Error
}

func (conn *itemsConn) UpdateQty(query interface{}, qtyValue int64) error {
	return conn.db.Model(&model.Inventory{}).Where(query).Update("quantity", qtyValue).Error
}

func UpdateQtyBasedOnBilling(tx *gorm.DB, productId uint, orderedQuantity int64) error {
	respData := new(model.Inventory)
	err := tx.Model(&model.Inventory{}).Where("ID = ?", productId).First(respData).Error
	if err != nil {
		return err
	}

	qty := respData.Quantity - orderedQuantity

	return tx.Model(&model.Inventory{}).Where("ID = ?", productId).Update("quantity", qty).Error
}
