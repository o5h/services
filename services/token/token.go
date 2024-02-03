package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	tokenExpirationTimeout = time.Second * 30
)

func Create(username string) (string, error) {

	claims := &AccessClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(tokenExpirationTimeout)}}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Revoke(tokenString string) {
	revoked.Revoke(tokenString, tokenExpirationTimeout)
}

func IsRevoked(tokenString string) bool {
	return revoked.IsRevoked(tokenString)
}

// ConvertAccessTokenToSubjectToken converts an access token to a subject token
func ConvertAccessTokenToSubjectToken(accessToken string) (string, error) {
	claims := AccessClaims{}
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})
	if err != nil {
		return "", fmt.Errorf("error parsing access token: %v", err)
	}

	// Create a new subject token
	expirationTime := time.Now().Add(tokenExpirationTimeout)
	subjectTokenClaims := &SubjectClaims{
		Username:         claims.Username,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)}}

	subjectToken := jwt.NewWithClaims(jwt.SigningMethodHS256, subjectTokenClaims)
	subjectTokenString, err := subjectToken.SignedString(SUBJECT_TOKEN_KEY)
	if err != nil {
		return "", fmt.Errorf("error creating subject token: %v", err)
	}

	return subjectTokenString, nil
}

func ConvertSubjectTokenToAccessToken(subjectToken string) (string, error) {
	// Parse and verify the subject token
	claims := &SubjectClaims{}
	token, err := jwt.ParseWithClaims(subjectToken, claims, func(token *jwt.Token) (interface{}, error) {
		return SUBJECT_TOKEN_KEY, nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing subject token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return "", fmt.Errorf("invalid subject token")
	}

	// Extract user information from the subject token
	username := claims.Username

	// Create a new access token
	accessTokenClaims := &AccessClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Set an expiration time for the access token
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(JWT_KEY)
	if err != nil {
		return "", fmt.Errorf("error creating access token: %v", err)
	}

	return accessTokenString, nil
}
