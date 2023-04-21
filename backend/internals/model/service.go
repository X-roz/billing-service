package model

import (
	"gorm.io/gorm"
)

type ServiceBilling struct {
	gorm.Model
	CustomerId     int
	Customer       Customer `gorm:"foreignKey:CustomerId; references:ID"`
	ServiceDetails []ServiceDetails
	Total          float64
}

type ServiceDetails struct {
	ID          uint    `gorm:"primarykey"`
	Description string  `json:"Description"`
	Price       float64 `json:"price"`
}
