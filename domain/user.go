package domain

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"not null" json:"name"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	Invoices []Invoice `gorm:"foreignKey:UserID" json:"invoices"`
	Events   []Event   `gorm:"foreignKey:UserID" json:"events"`
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

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
	SignUp(ctx context.Context, name, email, password string) error
}

type UserUseCase interface {
	SignIn(ctx context.Context, email, password string) (string, error)
	SignUp(ctx context.Context, name, email, password string) error
	Profile(ctx context.Context, email string) (User, error)
}
