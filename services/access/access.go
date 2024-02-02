package access

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/o5h/services/services/users"
	"golang.org/x/crypto/bcrypt"
)

type API interface {
	AuthenticateUser(username, password string) bool
	CreateToken(string, time.Duration) (string, error)
}

type service struct {
	userService users.API
}

var instance = &service{userService: users.GetService()}

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
	tokenString, err := token.SignedString(jwtKey)
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
