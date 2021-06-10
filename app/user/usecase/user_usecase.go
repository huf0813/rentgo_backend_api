package usecase

import (
	"context"
	"fmt"
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/infra/auth"
	"github.com/huf0813/rentgo_backend_api/utils/custom_security"
	"github.com/huf0813/rentgo_backend_api/utils/custom_storage"
	"mime/multipart"
	"time"
)

type UserUseCase struct {
	userRepoMysql domain.UserRepository
	timeOut       time.Duration
}

func NewUserUseCase(u domain.UserRepository, timeOut time.Duration) domain.UserUseCase {
	return &UserUseCase{
		userRepoMysql: u,
		timeOut:       timeOut,
	}
}

func (u *UserUseCase) SignIn(ctx context.Context, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	result, err := u.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := custom_security.NewValidatingValue(result.Password, password); err != nil {
		return "", err
	}

	duration := (24 * 30) * time.Hour
	token, err := auth.NewJWT(duration, email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUseCase) SignUp(ctx context.Context, name, email, password string) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	password, err := custom_security.NewHashingValue(password)
	if err != nil {
		return err
	}

	if err := u.userRepoMysql.SignUp(ctx, name, email, password); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) Profile(ctx context.Context, email string) (domain.UserProfileResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	result, err := u.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.UserProfileResponse{}, err
	}
	isVerified, err := u.userRepoMysql.CheckVerification(ctx, result.Email)
	if err != nil {
		return domain.UserProfileResponse{}, err
	}
	res := domain.UserProfileResponse{
		Email:      result.Email,
		Name:       result.Name,
		IsVerified: isVerified,
	}

	return res, nil
}

func (u *UserUseCase) UploadVerification(ctx context.Context,
	identityNumber,
	identityType string,
	identityImage *multipart.FileHeader,
	email string) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	path := fmt.Sprintf("%s/%s/%s", "assets", "image", "identity")
	filename, err := custom_storage.NewFileUpload(path, identityImage)
	if err != nil {
		return err
	}

	if err := u.userRepoMysql.UploadVerification(ctx,
		identityNumber,
		identityType,
		filename,
		email); err != nil {
		return err
	}

	return nil
}
