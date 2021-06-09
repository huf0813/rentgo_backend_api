package routes

import (
	"github.com/go-playground/validator"
	_cartHandler "github.com/huf0813/rentgo_backend_api/app/cart/delivery/http"
	_cartRepoMysql "github.com/huf0813/rentgo_backend_api/app/cart/repository/mysql"
	_cartUseCase "github.com/huf0813/rentgo_backend_api/app/cart/usecase"
	_productHandler "github.com/huf0813/rentgo_backend_api/app/product/delivery/http"
	_productRepository "github.com/huf0813/rentgo_backend_api/app/product/repository/mysql"
	_productUseCase "github.com/huf0813/rentgo_backend_api/app/product/usecase"
	_userHandler "github.com/huf0813/rentgo_backend_api/app/user/delivery/http"
	_userRepoMysql "github.com/huf0813/rentgo_backend_api/app/user/repository/mysql"
	_userUseCase "github.com/huf0813/rentgo_backend_api/app/user/usecase"
	_migrationHandler "github.com/huf0813/rentgo_backend_api/migration/delivery/http"
	_migrationRepoMysql "github.com/huf0813/rentgo_backend_api/migration/repository/mysql"
	_migrationUseCase "github.com/huf0813/rentgo_backend_api/migration/usecase"
	"github.com/huf0813/rentgo_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewRoutes(e *echo.Echo,
	db *gorm.DB,
	timeOut time.Duration,
	authMiddleware middleware.JWTConfig) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
			true,
			"welcome to API, please contact the administrator",
			nil))
	})

	userGroup := e.Group("/user", middleware.JWTWithConfig(authMiddleware))

	migrationRepoMysql := _migrationRepoMysql.NewMigrationRepoMysql(db)
	migrationUseCase := _migrationUseCase.NewMigrationUseCase(migrationRepoMysql, timeOut)
	_migrationHandler.NewMigrationHandler(e, migrationUseCase)

	userRepoMysql := _userRepoMysql.NewUserRepoMysql(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepoMysql, timeOut)
	_userHandler.NewUserHandler(e, userGroup, userUseCase)

	cartRepoMysql := _cartRepoMysql.NewCartRepoMysql(db)
	cartUseCase := _cartUseCase.NewCartUseCase(cartRepoMysql,
		userRepoMysql,
		timeOut)
	_cartHandler.NewCartHandler(userGroup, cartUseCase)

	productRepository := _productRepository.NewProductRepository(db)
	productUseCase := _productUseCase.NewProductUseCase(productRepository, timeOut)
	_productHandler.NewProductHandler(e, productUseCase)
}
