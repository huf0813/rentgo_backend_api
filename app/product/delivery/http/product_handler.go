package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/infra/auth"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductUseCase domain.ProductUseCase
}

func NewProductHandler(e *echo.Echo,
	userGroup *echo.Group,
	p domain.ProductUseCase) {
	handler := &ProductHandler{ProductUseCase: p}
	e.GET("/product", handler.SearchProduct)
	e.GET("/product/latest", handler.FetchProductLatest)
	e.GET("/product/trending", handler.FetchProductTrending)
	e.GET("/product/detail/:id", handler.FetchByID)
	e.GET("/product/detail/:id/reviews", handler.FetchReviewsByID)
	e.GET("/product/detail/:id/images", handler.FetchImagesByID)
	e.GET("/product/category/:category", handler.FetchByCategory)
	userGroup.GET("/vendor/product_category", handler.VendorProductCategory)
	userGroup.POST("/vendor/product/create", handler.VendorCreateProduct)
}

func (p *ProductHandler) VendorProductCategory(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := p.ProductUseCase.VendorFetchProductCategory(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(true,
		"get product category successfully",
		res))
}

func (p *ProductHandler) VendorCreateProduct(c echo.Context) error {
	productName := c.FormValue("product_name")
	productOverview := c.FormValue("product_overview")
	productCategory := c.FormValue("product_category")
	productCategoryInteger, err := strconv.Atoi(productCategory)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	productStock := c.FormValue("product_stock")
	productStockInteger, err := strconv.Atoi(productStock)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	productPrice := c.FormValue("product_price")
	productPriceInteger, err := strconv.Atoi(productPrice)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	productImage, err := c.FormFile("product_image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}
	pr := domain.ProductRequest{
		Name:            productName,
		Overview:        productOverview,
		ProductCategory: uint(productCategoryInteger),
		Stock:           uint(productStockInteger),
		Price:           uint(productPriceInteger),
		ProductImage:    productImage,
	}

	authorization := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(authorization)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := c.Request().Context()
	if err := p.ProductUseCase.VendorCreateProduct(ctx, &pr, token.Email); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"vendor create product successfully",
			nil))
}

func (p *ProductHandler) FetchProductTrending(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := p.ProductUseCase.FetchTrendingProduct(ctx)
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
			"get trending product successfully",
			res))
}

func (p *ProductHandler) FetchProductLatest(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := p.ProductUseCase.FetchLatestProduct(ctx)
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
			"get latest product successfully",
			res))
}

func (p *ProductHandler) FetchImagesByID(c echo.Context) error {
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
	res, err := p.ProductUseCase.FetchImagesByID(ctx, idInteger)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil))
	}

	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"get product images by id successfully",
			res))
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
				nil))
	}

	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"get product by id successfully",
			res))
}

func (p *ProductHandler) FetchReviewsByID(c echo.Context) error {
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
	res, err := p.ProductUseCase.FetchReviewsByID(ctx, idInteger)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil))
	}

	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"get product reviews by id successfully",
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
