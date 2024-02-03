package access

import (
	"encoding/base64"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/o5h/services/services/users"
	"golang.org/x/crypto/bcrypt"
)

type API interface {
	AuthenticateUser(username, password string) bool
	CreateToken(username string, timeout time.Duration) (string, error)
	InvalidateToken(token string)
	IsTokenValid(token string) bool
}

type service struct {
	userService         users.API
	tokenBlacklistMutex sync.RWMutex
	tokenBlacklist      map[string]struct{}
}

var (
	tokenExpirationTimeout = time.Second * 30
	instance               = &service{
		userService:    users.GetService(),
		tokenBlacklist: make(map[string]struct{}),
	}
)

func GetService() API {
	return instance
}

func (serv *service) CreateToken(username string, timeout time.Duration) (string, error) {
	expirationTime := time.Now().Add(timeout)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (serv *service) AuthenticateUser(username, password string) bool {
	user, err := serv.userService.GetUser(username)
	if err != nil {
		log.Println("Unknown user", username)
		return false
	}
	hashedPassword, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		log.Println("Corrupted password for user", username)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil && username == user.Username
}

func (serv *service) InvalidateToken(tokenString string) {
	serv.tokenBlacklistMutex.Lock()
	defer serv.tokenBlacklistMutex.Unlock()
	time.AfterFunc(tokenExpirationTimeout, func() {
		serv.RemoveToken(tokenString)
	})
	serv.tokenBlacklist[tokenString] = struct{}{}
}

func (serv *service) IsTokenValid(tokenString string) bool {
	serv.tokenBlacklistMutex.RLock()
	defer serv.tokenBlacklistMutex.RUnlock()
	_, exists := serv.tokenBlacklist[tokenString]
	return !exists
}

func (serv *service) RemoveToken(tokenString string) {
	serv.tokenBlacklistMutex.Lock()
	defer serv.tokenBlacklistMutex.Unlock()
	delete(serv.tokenBlacklist, tokenString)
	log.Println("Token removed", tokenString)
}

// ConvertAccessTokenToSubjectToken converts an access token to a subject token
func (serv *service) ConvertAccessTokenToSubjectToken(accessToken string) (string, error) {
	claims := Claims{}
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})

	// Create a new subject token
	expirationTime := time.Now().Add(tokenExpirationTimeout)
	subjectTokenClaims := &SubjectTokenClaims{
		Username: claims.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix()}}

	subjectToken := jwt.NewWithClaims(jwt.SigningMethodHS256, subjectTokenClaims)
	subjectTokenString, err := subjectToken.SignedString(SUBJECT_TOKEN_KEY)
	if err != nil {
		return "", fmt.Errorf("error creating subject token: %v", err)
	}

	return subjectTokenString, nil
}
