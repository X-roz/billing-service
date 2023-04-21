package model

import (
	"time"
)

type SalesBilling struct {
	ID           uint `gorm:"primarykey"`
	CustomerId   uint
	Customer     Customer       `gorm:"foreignKey:CustomerId; references:ID"`
	SalesDetails []SalesDetails `gorm:"foreignkey:OrderId; references:ID"`
	SummaryTotal float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SalesDetails struct {
	ID              uint `gorm:"primarykey"`
	OrderId         uint
	ProductId       uint
	Product         Inventory `gorm:"foreignKey:ProductId; references:ID"`
	OrderedQuantity int64
	Total           float64
}
