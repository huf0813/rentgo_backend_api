package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CartHandler struct {
	CartUseCase domain.CartUseCase
}

func NewCartHandler(e *echo.Echo, c domain.CartUseCase) {
	handler := &CartHandler{CartUseCase: c}
	e.POST("/checkout", handler.Checkout)
}

func (c *CartHandler) Checkout(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true, "", nil))
}
