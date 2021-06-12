package mysql

import (
	"context"
	"errors"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
)

type UserRepoMysql struct {
	DB *gorm.DB
}

func NewUserRepoMysql(db *gorm.DB) domain.UserRepository {
	return &UserRepoMysql{DB: db}
}

func (u *UserRepoMysql) SearchVendor(ctx context.Context, storeName string) error {
	return nil
}

func (u *UserRepoMysql) CheckVerification(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := u.DB.
		WithContext(ctx).
		Model(&domain.User{}).
		Where("users.email = ?", email).
		Where("users.identity_number IS NOT NULL").
		Where("users.identity_type IS NOT NULL").
		Where("users.identity_image IS NOT NULL").
		Count(&count).Error; err != nil {
		return false, err
	}
	if count <= 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (u *UserRepoMysql) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	result := u.DB.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user)
	if err := result.Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepoMysql) SignUp(ctx context.Context, name, email, password string) error {
	user := domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	result := u.DB.WithContext(ctx).Create(&user)
	if err := result.Error; err != nil {
		return err
	}
	if rows := result.RowsAffected; rows <= 0 {
		return errors.New("failed to insert data, empty feedback")
	}
	return nil
}

func (u *UserRepoMysql) UploadVerification(ctx context.Context,
	identityNumber,
	identityType,
	identityImage,
	storeName,
	storePhone,
	email string) error {
	updateIdentity := domain.User{
		IdentityNumber: identityNumber,
		IdentityType:   identityType,
		IdentityImage:  identityImage,
		StoreName:      storeName,
		Phone:          storePhone,
	}

	if err := u.DB.
		WithContext(ctx).
		Model(&domain.User{}).
		Where("users.email = ?", email).
		Updates(&updateIdentity).Error; err != nil {
		return err
	}

	return nil
}
