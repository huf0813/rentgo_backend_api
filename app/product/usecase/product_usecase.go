package usecase

import (
	"context"
	"errors"
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/utils/custom_storage"
	"time"
)

type ProductUseCase struct {
	productRepoMysql      domain.ProductRepository
	ProductImageRepoMysql domain.ProductImageRepository
	userRepoMysql         domain.UserRepository
	timeOut               time.Duration
}

func NewProductUseCase(p domain.ProductRepository,
	u domain.UserRepository,
	pi domain.ProductImageRepository,
	timeOut time.Duration) domain.ProductUseCase {
	return &ProductUseCase{
		productRepoMysql:      p,
		ProductImageRepoMysql: pi,
		userRepoMysql:         u,
		timeOut:               timeOut,
	}
}

func (p *ProductUseCase) VendorFetchProductCategory(ctx context.Context) ([]domain.ProductCategoryResponse,
	error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res, err := p.productRepoMysql.VendorFetchProductCategory(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *ProductUseCase) VendorCreateProduct(ctx context.Context,
	pr *domain.ProductRequest,
	email string) error {
	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	isVerified, err := p.userRepoMysql.CheckVerification(ctx, email)
	if err != nil {
		return err
	}
	if !isVerified {
		return errors.New("account is not verified, try again later")
	}

	user, err := p.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	createProduct := domain.Product{
		Name:              pr.Name,
		Overview:          pr.Overview,
		Price:             pr.Price,
		Stock:             pr.Stock,
		ProductCategoryID: pr.ProductCategory,
		UserID:            user.ID,
	}
	productID, err := p.productRepoMysql.VendorCreateProduct(ctx, &createProduct)
	if err != nil {
		return err
	}

	filename, err := custom_storage.NewFileUpload("assets/image/product", pr.ProductImage)
	if err != nil {
		return err
	}

	if err := p.ProductImageRepoMysql.CreateProductImage(ctx, productID, filename); err != nil {
		return err
	}

	return nil
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
