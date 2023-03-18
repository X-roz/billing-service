package items

import (
	"billing-service/internals/model"
	"billing-service/internals/repository"
)

type Service interface {
	AddItems(*model.ItemDetails) error
	GetItems(interface{}) (*[]model.ItemDetails, error)
	UpdateItem(interface{}, *model.ItemDetails) error
	UpdateQty(interface{}, int64) error
}

type service struct {
	ItemsRepository repository.ItemsRepository
}

func NewService() *service {
	return &service{repository.NewItemsConn()}
}

func (s *service) AddItems(data *model.ItemDetails) error {
	return s.ItemsRepository.AddItems(data)
}

func (s *service) GetItems(query interface{}) (*[]model.ItemDetails, error) {
	return s.ItemsRepository.GetItems(query)
}

func (s *service) UpdateItem(query interface{}, data *model.ItemDetails) error {
	return s.ItemsRepository.UpdateItem(query, data)
}

func (s *service) UpdateQty(query interface{}, qtyValue int64) error {
	return s.ItemsRepository.UpdateQty(query, qtyValue)
}
