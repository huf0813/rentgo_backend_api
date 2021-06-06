package auth

import "github.com/dgrijalva/jwt-go"

type CustomToken struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
