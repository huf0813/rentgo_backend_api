package domain

import (
	"context"
	"gorm.io/gorm"
	"mime/multipart"
)

type User struct {
	gorm.Model
	Name           string    `gorm:"not null" json:"name"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Password       string    `gorm:"not null" json:"password"`
	Phone          string    `gorm:"unique;default:null" json:"phone"`
	StoreName      string    `gorm:"unique;default:null" json:"store_name"`
	IdentityType   string    `gorm:"default:null" json:"identity_type"`
	IdentityNumber string    `gorm:"unique;default:null" json:"identity_number"`
	IdentityImage  string    `gorm:"unique;default:null" json:"identity_image"`
	Invoices       []Invoice `gorm:"foreignKey:UserID" json:"invoices"`
	Events         []Event   `gorm:"foreignKey:UserID" json:"events"`
	Carts          []Cart    `gorm:"foreignKey:UserID" json:"carts"`
	Products       []Product `gorm:"foreignKey:UserID" json:"products"`
}

type UserSignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserSignUpRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UserProfileResponse struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
}

type VendorResponse struct {
	VendorID           uint   `json:"vendor_id"`
	VendorPhone        string `json:"vendor_phone"`
	VendorName         string `json:"vendor_name"`
	VendorAddress      string `json:"vendor_address"`
	VendorEmail        string `json:"vendor_email"`
	VendorCountReviews string `json:"vendor_count_reviews"`
	VendorRating       string `json:"vendor_rating"`
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
	SignUp(ctx context.Context, name, email, password string) error
	UploadVerification(ctx context.Context,
		identityNumber,
		identityType,
		identityImage,
		StoreName,
		StorePhone,
		email string) error
	CheckVerification(ctx context.Context, email string) (bool, error)
	SearchVendor(ctx context.Context, storeName string) error
}

type UserUseCase interface {
	SignIn(ctx context.Context, email, password string) (string, error)
	SignUp(ctx context.Context, name, email, password string) error
	Profile(ctx context.Context, email string) (UserProfileResponse, error)
	UploadVerification(ctx context.Context,
		identityNumber,
		identityType string,
		identityImage *multipart.FileHeader,
		StoreName,
		StorePhone,
		email string) error
}
