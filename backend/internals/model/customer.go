package model

type Customer struct {
	ID           uint   `gorm:"primarykey"`
	Name         string `json:"CustomerName"`
	MobileNumber string `json:"CustomerPhoneNumber" gorm:"uniqueIndex"`
	GstId        string
}
