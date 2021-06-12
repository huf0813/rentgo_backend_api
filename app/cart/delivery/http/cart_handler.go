package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/infra/auth"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CartHandler struct {
	CartUseCase domain.CartUseCase
}

func NewCartHandler(userGroup *echo.Group, c domain.CartUseCase) {
	handler := &CartHandler{CartUseCase: c}
	userGroup.POST("/cart/create/:product_id", handler.AddProductToCart)
	userGroup.DELETE("/cart/delete/:cart_id", handler.DeleteProductToCart)
	userGroup.GET("/cart", handler.FetchCart)
}

func (c *CartHandler) DeleteProductToCart(echoContext echo.Context) error {
	cartID := echoContext.Param("cart_id")
	cartIDInteger, err := strconv.Atoi(cartID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	bearer := echoContext.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := echoContext.Request().Context()
	if err := c.CartUseCase.DeleteCartByID(ctx,
		token.Email,
		uint(cartIDInteger)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return echoContext.JSON(http.StatusOK, custom_response.NewCustomResponse(
		false,
		"delete product from cart successfully",
		nil))
}

func (c *CartHandler) AddProductToCart(echoContext echo.Context) error {
	productID := echoContext.Param("product_id")
	productIDInteger, err := strconv.Atoi(productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	bearer := echoContext.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	addCartRequest := new(domain.CartAddProductRequest)
	if err := echoContext.Bind(addCartRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	if err := echoContext.Validate(addCartRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := echoContext.Request().Context()
	if err := c.CartUseCase.AddProductToCart(ctx,
		productIDInteger,
		token.Email,
		addCartRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return echoContext.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"add to cart successfully",
		nil))
}

func (c *CartHandler) FetchCart(echoContext echo.Context) error {
	bearer := echoContext.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := echoContext.Request().Context()
	result, err := c.CartUseCase.FetchCart(ctx, token.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return echoContext.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"fetch cart successfully",
		result))
}
