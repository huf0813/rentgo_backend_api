package mysql

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
)

type CartRepoMysql struct {
	DB *gorm.DB
}

func NewCartRepoMysql(db *gorm.DB) domain.CartRepository {
	return &CartRepoMysql{DB: db}
}

func (c *CartRepoMysql) AddProductToCart(ctx context.Context,
	quantity int,
	productID,
	userID uint) error {
	newCart := domain.Cart{
		ProductID: productID,
		UserID:    userID,
		Quantity:  uint(quantity),
	}

	if err := c.DB.
		WithContext(ctx).
		Create(&newCart).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartRepoMysql) FetchCart(ctx context.Context, userID uint) ([]domain.CartResponse, error) {
	var result []domain.CartResponse

	rows, err := c.DB.WithContext(ctx).
		Raw("select p.name as product_name, "+
			"vendor.name as product_vendor, "+
			"c.quantity as product_quantity, "+
			"c.id as cart_id, "+
			"p.price as product_price from products p "+
			"JOIN carts c on p.id = c.product_id "+
			"JOIN users user_cart on c.user_id = user_cart.id "+
			"JOIN users vendor on p.user_id = vendor.id "+
			"where user_cart.id = ?", userID).
		Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var row domain.CartResponse
		if err := rows.Scan(&row.ProductName, &row.Vendor, &row.Quantity, &row.ID, &row.ProductPrice); err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}

func (c *CartRepoMysql) FetchCartByID(ctx context.Context, userID, cartID uint) (domain.Cart, error) {
	var row domain.Cart

	if err := c.DB.
		WithContext(ctx).
		Model(&domain.Cart{}).
		Where("carts.id = ?", cartID).
		First(&row, userID).Error; err != nil {
		return domain.Cart{}, err
	}

	return row, nil
}
