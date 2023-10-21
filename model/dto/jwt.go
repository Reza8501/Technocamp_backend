package dto

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_EXPIRATION_DURATION = time.Duration(1) * time.Minute

type JwtClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type ResponseJwt struct {
	Token string `json:"token"`
}
