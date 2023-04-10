package route

import (
	inventory "billing-service/internals/app/inventory"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	inventory.NewHandler().Route(g.Group("inventory"))
}
