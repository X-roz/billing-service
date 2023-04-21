package sales

import (
	"billing-service/response"
	"time"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler() *handler {
	return &handler{NewService()}
}

func (h *handler) AddBilling(c echo.Context) error {

	s := new(SalesPayload)
	c.Bind(s)
	err := h.service.AddSalesBilling(s)
	if err != nil {
		return response.RespErr(c, "Add Items error", err)
	}

	return response.RespSuccess(c, "Item added successfully")
}

func (h *handler) SalesBillings(c echo.Context) error {

	var startdate time.Time
	var enddate time.Time

	period := c.QueryParam("period")
	switch period {
	case "last-month":
		startdate = time.Now().AddDate(0, -1, 0)
		enddate = time.Now()
	case "last-week":
		startdate = time.Now().AddDate(0, 0, -7)
		enddate = time.Now()
	default:
		startdate = time.Now().AddDate(0, 0, -1)
		enddate = time.Now()
	}

	data, err := h.service.AllSalesBillingDetails(startdate, enddate)
	if err != nil {
		return response.RespErr(c, "data retrieving error", err)
	}
	return response.RespSuccessInfo(c, "Retrieved successfully", data)
}
