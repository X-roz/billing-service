package servicebilling

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.POST("/add", h.AddBilling)
	g.GET("/get", h.ServiceBillings)
}
