package route

import (
	items "billing-service/internals/app/Items"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	items.NewHandler().Route(g)
}
