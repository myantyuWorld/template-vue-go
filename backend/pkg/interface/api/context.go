package api

import (
	"context"

	"github.com/labstack/echo/v4"
)

func NewContext(c echo.Context) context.Context {
	return c.Request().Context()
}
