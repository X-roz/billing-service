package model

type ItemDetails struct {
	ID uint `gorm:"primarykey"`
	// ItemId   int32
	Name     string  `json:"name"`
	Tax      float64 `json:"tax"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}
