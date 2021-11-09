package api

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func RegisterRouters(e *echo.Echo) {
	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "service up",
		})
	})
}
