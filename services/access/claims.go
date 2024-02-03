package access

import (
	"github.com/golang-jwt/jwt/v5"
)

// AccessTokenClaims struct represents the Access JWT claims
type AccessTokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// AccessTokenClaims struct represents the Subject JWT claims
type SubjectTokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
