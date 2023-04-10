package inventory

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {

	g.POST("/add-item", h.AddItems)
	g.GET("/get-items", h.RetrieveItems)
	g.POST("/update-item", h.UpdateItem)
	// g.GET("/update-qty", h.UpdateQty)
	g.POST("/update-bulk", h.UpdateBulk)

}
