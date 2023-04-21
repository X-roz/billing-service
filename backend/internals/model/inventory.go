package model

type Inventory struct {
	ID uint `gorm:"primarykey"`
	// ItemId   int32
	Name     string  `json:"name" gorm:"uniqueIndex"`
	Tax      float64 `json:"tax"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}
