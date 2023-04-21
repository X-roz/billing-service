package sales

type SalesPayload struct {
	CustomerName        string
	CustomerPhoneNumber string
	SalesDetails        []Details
	SummaryTotal        float64
}

type Details struct {
	ProductId       uint
	OrderedQuantity int64
	Total           float64
}
