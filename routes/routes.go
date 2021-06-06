package routes

import (
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

func NewRoutes(e *echo.Echo,
	db *gorm.DB,
	timeOut time.Duration,
	authMiddleware middleware.JWTConfig) {
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

	productRepository := _productRepository.NewProductRepository(db)
	productUseCase := _productUseCase.NewProductUseCase(productRepository, timeOut)
	_productHandler.NewProductHandler(e, productUseCase)
}
