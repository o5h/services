package access

import "github.com/golang-jwt/jwt"

// AccessTokenClaims struct represents the Access JWT claims
type AccessTokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// AccessTokenClaims struct represents the Subject JWT claims
type SubjectTokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
