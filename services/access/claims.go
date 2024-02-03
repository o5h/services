package access

import "github.com/golang-jwt/jwt"

// Claims struct represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type SubjectTokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	// Add other claims as needed
	jwt.StandardClaims
}
