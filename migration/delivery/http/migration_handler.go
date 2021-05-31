package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MigrationHandler struct {
	MigrationUseCase domain.MigrationUseCase
}

func NewMigrationHandler(e *echo.Echo, m domain.MigrationUseCase) {
	handler := &MigrationHandler{MigrationUseCase: m}
	e.GET("/migrate", handler.Migrate)
}

func (m *MigrationHandler) Migrate(c echo.Context) error {
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"migrated successfully",
		nil))
}
