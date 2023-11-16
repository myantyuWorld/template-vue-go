package middleware

import "github.com/labstack/echo/v4"

func MakeErrorHandlerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				// do something
			}

			return nil
		}
	}
}
