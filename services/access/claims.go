package access

import "github.com/golang-jwt/jwt"

// Claims struct represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
