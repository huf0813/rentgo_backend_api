package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductUseCase domain.ProductUseCase
}

func NewProductHandler(e *echo.Echo, p domain.ProductUseCase) {
	handler := &ProductHandler{ProductUseCase: p}
	e.GET("/product", handler.SearchProduct)
	e.GET("/product/:id", handler.FetchByID)
	e.GET("/product/:category", handler.FetchByCategory)
}

func (p *ProductHandler) SearchProduct(c echo.Context) error {
	ctx := c.Request().Context()
	name := c.QueryParam("name")
	res, err := p.ProductUseCase.SearchProduct(ctx, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				res))
	}

	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"search product by name successfully",
			res))
}

func (p *ProductHandler) FetchByID(c echo.Context) error {
	id := c.Param("id")
	idInteger, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil))
	}

	ctx := c.Request().Context()
	res, err := p.ProductUseCase.FetchByID(ctx, idInteger)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				res))
	}

	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"get product by id successfully",
			res))
}

func (p *ProductHandler) FetchByCategory(c echo.Context) error {
	category := c.Param("category")
	ctx := c.Request().Context()
	res, err := p.ProductUseCase.FetchByCategory(ctx, category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				res))
	}

	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"get product by category successfully",
			res))
}
