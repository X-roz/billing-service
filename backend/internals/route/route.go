package route

import (
	inventory "billing-service/internals/app/inventory"
	"billing-service/internals/app/sales"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	inventory.NewHandler().Route(g.Group("inventory"))
	sales.NewHandler().Route(g.Group("sales-billing"))
}
