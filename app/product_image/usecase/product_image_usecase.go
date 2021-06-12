package usecase

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"time"
)

type ProductImageUseCase struct {
	productImageRepoMysql domain.ProductImageRepository
	timeOut               time.Duration
}
