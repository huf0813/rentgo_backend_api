package routes

import (
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func NewRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
			true,
			"welcome to API, please contact the administrator",
			nil))
	})
}
