package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAPIServer() *APIServer {
	e := echo.New()
	e.Use(middleware.Recover())

	return &APIServer{e}
}

type APIServer struct {
	*echo.Echo
}
