package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/infra/auth"
	"github.com/huf0813/rentgo_backend_api/utils/custom_converter"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type InvoiceHandler struct {
	InvoiceUseCase domain.InvoiceUseCase
}

func NewInvoiceHandler(userGroup *echo.Group, i domain.InvoiceUseCase) {
	handler := &InvoiceHandler{InvoiceUseCase: i}
	userGroup.POST("/checkout",
		handler.Checkout)
	userGroup.PUT("/on_going/:receipt_number",
		handler.OnGoing)
	userGroup.PUT("/completed/:receipt_number",
		handler.Completed)
	userGroup.PUT("/review/create/:invoice_product_id",
		handler.CreateReviewByInvoiceProductID)
	userGroup.GET("/invoice/accepted",
		handler.GetInvoicesAccepted)
	userGroup.GET("/invoice/on_going",
		handler.GetInvoicesOnGoing)
	userGroup.GET("/invoice/product/:receipt_code",
		handler.GetInvoiceProducts)
	userGroup.GET("/invoice/completed",
		handler.GetInvoicesCompleted)
	userGroup.GET("/invoice/completed/product/:receipt_code",
		handler.GetCompletedInvoiceProducts)
}

func (i *InvoiceHandler) CreateReviewByInvoiceProductID(c echo.Context) error {
	createReview := new(domain.InvoiceReviewRequest)
	if err := c.Bind(createReview); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(createReview); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	invoiceProductID := c.Param("invoice_product_id")
	invoiceProductIDInteger, err := strconv.Atoi(invoiceProductID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil))
	}

	if err := i.InvoiceUseCase.CreateReviewByInvoiceProductID(ctx,
		uint(invoiceProductIDInteger),
		createReview.Rating,
		createReview.Review); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil))
	}

	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			false,
			"create review successfully",
			nil))
}

func (i *InvoiceHandler) GetCompletedInvoiceProducts(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	receiptCode := c.Param("receipt_code")
	ctx := c.Request().Context()
	res, err := i.InvoiceUseCase.GetCompletedInvoiceProductsByReceiptCode(ctx,
		token.Email,
		receiptCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(true,
		"fetch completed product successfully",
		res))
}

func (i *InvoiceHandler) GetInvoiceProducts(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	receiptCode := c.Param("receipt_code")
	ctx := c.Request().Context()
	res, err := i.InvoiceUseCase.GetInvoiceProductsByReceiptCode(ctx,
		token.Email,
		receiptCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"get invoice products successfully",
		res))
}

func (i *InvoiceHandler) GetInvoicesAccepted(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := c.Request().Context()
	res, err := i.InvoiceUseCase.GetInvoicesAccepted(ctx, token.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"get invoices accepted successfully",
		res))
}

func (i *InvoiceHandler) GetInvoicesOnGoing(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := c.Request().Context()
	res, err := i.InvoiceUseCase.GetInvoicesOnGoing(ctx, token.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"get invoices on going successfully",
		res))
}

func (i *InvoiceHandler) GetInvoicesCompleted(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	ctx := c.Request().Context()
	res, err := i.InvoiceUseCase.GetInvoicesCompleted(ctx,
		token.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"get invoices completed successfully",
		res))
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
	if err := i.InvoiceUseCase.UpdateInvoiceOnGoing(ctx,
		token.Email,
		invoiceID); err != nil {
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

func (i *InvoiceHandler) Completed(c echo.Context) error {
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
	if err := i.InvoiceUseCase.UpdateInvoiceCompleted(ctx,
		token.Email,
		invoiceID); err != nil {
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
