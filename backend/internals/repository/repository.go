package repository

import (
	"billing-service/db"
	"billing-service/internals/model"

	"gorm.io/gorm"
)

type ItemsRepository interface {
	AddItems(*model.ItemDetails) error
	GetItems(interface{}) (*[]model.ItemDetails, error)
	UpdateItem(interface{}, *model.ItemDetails) error
	UpdateQty(interface{}, int64) error
}

type itemsConn struct {
	db *gorm.DB
}

func NewItemsConn() *itemsConn {
	return &itemsConn{db.DbManager()}
}

func (conn *itemsConn) AddItems(data *model.ItemDetails) error {
	return conn.db.Model(&model.ItemDetails{}).Create(data).Error
}

func (conn *itemsConn) GetItems(query interface{}) (*[]model.ItemDetails, error) {
	respData := new([]model.ItemDetails)
	err := conn.db.Model(&model.ItemDetails{}).Where(query).Find(respData).Error
	return respData, err
}

func (conn *itemsConn) UpdateItem(query interface{}, data *model.ItemDetails) error {
	return conn.db.Model(&model.ItemDetails{}).Where(query).Updates(data).Error
}

func (conn *itemsConn) UpdateQty(query interface{}, qtyValue int64) error {
	return conn.db.Model(&model.ItemDetails{}).Where(query).Update("quantity", qtyValue).Error
}
