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
	e.GET("/seed", handler.Seed)
	e.GET("/faker", handler.Faker)
	e.GET("/drop", handler.Drop)
}

func (m *MigrationHandler) Migrate(c echo.Context) error {
	ctx := c.Request().Context()
	if err := m.MigrationUseCase.Migrate(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"migrated successfully",
		nil))
}

func (m *MigrationHandler) Seed(c echo.Context) error {
	ctx := c.Request().Context()
	if err := m.MigrationUseCase.Seed(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"seeded successfully",
		nil))
}

func (m *MigrationHandler) Faker(c echo.Context) error {
	ctx := c.Request().Context()
	if err := m.MigrationUseCase.Faker(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"faker generated successfully",
		nil))
}

func (m *MigrationHandler) Drop(c echo.Context) error {
	ctx := c.Request().Context()
	if err := m.MigrationUseCase.Drop(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"Drop tables successfully",
		nil))
}
