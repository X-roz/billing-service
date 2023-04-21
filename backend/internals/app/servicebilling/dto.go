package servicebilling

type ServicePayload struct {
	CustomerName        string
	CustomerPhoneNumber string
	ServiceDetails      []Details
	SummaryTotal        float64
}

type Details struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
