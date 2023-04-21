package model

import "time"

type ServiceBilling struct {
	ID             uint `gorm:"primarykey"`
	CustomerId     uint
	Customer       Customer         `gorm:"foreignKey:CustomerId; references:ID"`
	ServiceDetails []ServiceDetails `gorm:"foreignkey:OrderId; references:ID"`
	SummaryTotal   float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ServiceDetails struct {
	ID          uint `gorm:"primarykey"`
	OrderId     uint
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
