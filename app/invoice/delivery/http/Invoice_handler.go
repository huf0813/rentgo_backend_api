package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type InvoiceHandler struct {
	InvoiceUseCase domain.InvoiceUseCase
}

func NewInvoiceHandler(e *echo.Echo, i domain.InvoiceUseCase) {
	handler := &InvoiceHandler{InvoiceUseCase: i}
	e.POST("/", handler.Checkout)
}

func (i *InvoiceHandler) Checkout(c echo.Context) error {
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"",
		nil))
}
