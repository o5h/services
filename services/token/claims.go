package token

import (
	"github.com/golang-jwt/jwt/v5"
)

// AccessClaims struct represents the Access JWT claims
type AccessClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// AccessTokenClaims struct represents the Subject JWT claims
type SubjectClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
