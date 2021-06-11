package usecase

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"time"
)

type ProductUseCase struct {
	productRepoMysql domain.ProductRepository
	timeOut          time.Duration
}

func NewProductUseCase(p domain.ProductRepository, timeOut time.Duration) domain.ProductUseCase {
	return &ProductUseCase{
		productRepoMysql: p,
		timeOut:          timeOut,
	}
}

func (p *ProductUseCase) FetchTrendingProduct(ctx context.Context) ([]domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res, err := p.productRepoMysql.FetchTrendingProduct(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *ProductUseCase) FetchLatestProduct(ctx context.Context) ([]domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res, err := p.productRepoMysql.FetchLatestProduct(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *ProductUseCase) FetchImagesByID(ctx context.Context, id int) ([]domain.ProductImageResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res, err := p.productRepoMysql.FetchImagesByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *ProductUseCase) FetchByID(ctx context.Context, id int) (domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res, err := p.productRepoMysql.FetchByID(ctx, id)
	if err != nil {
		return domain.ProductResponse{}, err
	}

	return res, nil
}

func (p *ProductUseCase) FetchByCategory(ctx context.Context, category string) ([]domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res, err := p.productRepoMysql.FetchByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p ProductUseCase) SearchProduct(ctx context.Context, name string) ([]domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res, err := p.productRepoMysql.SearchProduct(ctx, name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *ProductUseCase) FetchReviewsByID(ctx context.Context, id int) ([]domain.ProductReviewResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res, err := p.productRepoMysql.FetchReviewsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
