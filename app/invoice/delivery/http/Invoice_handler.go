package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/infra/auth"
	"github.com/huf0813/rentgo_backend_api/utils/custom_converter"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InvoiceHandler struct {
	InvoiceUseCase domain.InvoiceUseCase
}

func NewInvoiceHandler(userGroup *echo.Group, i domain.InvoiceUseCase) {
	handler := &InvoiceHandler{InvoiceUseCase: i}
	userGroup.POST("/checkout", handler.Checkout)
	userGroup.PUT("/on_going/:receipt_number", handler.OnGoing)
	userGroup.PUT("/completed/:receipt_number", handler.Completed)
}

func (i *InvoiceHandler) Completed(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := c.Request().Context()
	invoiceID := c.Param("receipt_number")
	if err := i.InvoiceUseCase.UpdateInvoiceCompleted(ctx,
		invoiceID,
		token.Email); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"success update invoice to completed",
		nil))
}

func (i *InvoiceHandler) Checkout(c echo.Context) error {
	checkoutRequest := new(domain.InvoiceCheckoutRequest)
	if err := c.Bind(checkoutRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	if err := c.Validate(checkoutRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	startDate, err := custom_converter.NewDateConverter(checkoutRequest.StartDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	finishDate, err := custom_converter.NewDateConverter(checkoutRequest.FinishDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := c.Request().Context()
	if err := i.InvoiceUseCase.CreateCheckOut(ctx,
		startDate, finishDate,
		token.Email,
		checkoutRequest.CartIDS); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"checkout product successfully",
		nil))
}

func (i *InvoiceHandler) OnGoing(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	invoiceID := c.Param("receipt_number")
	ctx := c.Request().Context()
	if err := i.InvoiceUseCase.UpdateInvoiceOnGoing(ctx, token.Email, invoiceID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"success update invoice to on going",
		nil))
}
