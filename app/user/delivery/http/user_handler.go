package http

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/infra/auth"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(e *echo.Echo, userGroup *echo.Group, u domain.UserUseCase) {
	handler := &UserHandler{u}
	userGroup.GET("/profile", handler.Profile)
	userGroup.PUT("/profile/verification", handler.Verification)

	authGroup := e.Group("/auth")
	authGroup.POST("/sign_in", handler.SignIn)
	authGroup.POST("/sign_up", handler.SignUp)
}

func (u *UserHandler) SignIn(c echo.Context) error {
	user := new(domain.UserSignInRequest)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	result, err := u.UserUseCase.SignIn(ctx, user.Email, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.
			NewCustomResponse(false, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"sign in successfully",
		result),
	)
}

func (u *UserHandler) SignUp(c echo.Context) error {
	user := new(domain.UserSignUpRequest)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if err := u.UserUseCase.SignUp(ctx,
		user.Name,
		user.Email,
		user.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, custom_response.
			NewCustomResponse(false, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"sign up successfully",
		nil),
	)
}

func (u *UserHandler) Profile(c echo.Context) error {
	ctx := c.Request().Context()
	authorization := c.Request().Header.Get("Authorization")
	claims, err := auth.NewTokenExtraction(authorization)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := u.UserUseCase.Profile(ctx, claims.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"get user's profile successfully",
		result),
	)
}

func (u *UserHandler) Verification(c echo.Context) error {
	ctx := c.Request().Context()
	authorization := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(authorization)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	storePhone := c.FormValue("store_phone")
	storeName := c.FormValue("store_name")
	identityType := c.FormValue("identity_type")
	identityNumber := c.FormValue("identity_number")
	identityImage, err := c.FormFile("identity_image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	if err := u.UserUseCase.UploadVerification(ctx,
		identityNumber,
		identityType,
		identityImage,
		storeName,
		storePhone,
		token.Email); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_response.NewCustomResponse(
			false,
			err.Error(),
			nil))
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"upload user's identity successfully",
		nil),
	)
}
