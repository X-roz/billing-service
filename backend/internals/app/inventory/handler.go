package inventory

import (
	"billing-service/internals/model"
	"billing-service/response"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler() *handler {
	return &handler{NewService()}
}

func (h *handler) AddItems(c echo.Context) error {

	d := new(model.ItemDetails)
	c.Bind(d)

	query := map[string]string{"name": d.Name}
	idata, _ := h.service.GetItems(query)
	if len(*idata) > 0 {
		return response.RespErr(c, "Items already available", errors.New("please verify the item name"))
	}

	// rand.Seed(time.Now().UnixNano())
	// min := 1000
	// max := 10000
	// d.ItemId = rand.Intn(max-min+1) + min

	err := h.service.AddItems(d)
	if err != nil {
		return response.RespErr(c, "Add Items error", err)
	}

	return response.RespSuccess(c, "Item added successfully")
}

func (h *handler) RetrieveItems(c echo.Context) error {

	query := map[string]string{}
	respData, err := h.service.GetItems(query)
	if err != nil {
		return response.RespErr(c, "Retriving Items failed", err)
	}

	return response.RespSuccessInfo(c, "Item added successfully", respData)
}

func (h *handler) UpdateItem(c echo.Context) error {

	d := new(model.ItemDetails)
	c.Bind(d)

	query := map[string]string{"id": strconv.FormatUint(uint64(d.ID), 32)}
	ib, _ := h.service.GetItems(query)
	if len(*ib) <= 0 {
		return response.RespErr(c, "Items not found", errors.New("please verify the item name"))
	}

	err := h.service.UpdateItem(query, d)
	if err != nil {
		return response.RespErr(c, "Update Item failed", err)
	}

	return response.RespSuccess(c, "Item updated successfully")
}

func (h *handler) UpdateQty(c echo.Context) error {
	name := c.QueryParam("name")
	qty := c.QueryParam("quantity")

	query := map[string]string{"name": name}

	ib, _ := h.service.GetItems(query)
	if len(*ib) <= 0 {
		return response.RespErr(c, "Items not found", errors.New("please verify the item name"))
	}

	qtyValue, err := strconv.Atoi(qty)
	if err != nil {
		return response.RespErr(c, "Internal string to int conversion error", err)
	}

	err = h.service.UpdateQty(query, int64(qtyValue))
	if err != nil {
		return response.RespErr(c, "update failed", err)
	}

	return response.RespSuccess(c, "Item Qty updated successfully")
}

func (h *handler) UpdateBulk(c echo.Context) error {

	d := new([]model.ItemDetails)
	c.Bind(d)

	for _, i := range *d {
		query := map[string]string{"id": strconv.FormatUint(uint64(i.ID), 32)}
		ib, _ := h.service.GetItems(query)
		if len(*ib) <= 0 {
			return response.RespErr(c, "Items not found", errors.New("please verify the item name"))
		}

		err := h.service.UpdateItem(query, &i)
		if err != nil {
			return response.RespErr(c, "Update Item failed", err)
		}
	}

	return response.RespSuccess(c, "Item updated successfully")

}
