package user

import (
	"encoding/base64"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorUserAlreadyExists = errors.New("User already exists")
	ErrorUserNotFound      = errors.New("User not found")
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users map[string]User = make(map[string]User)

func Create(u User) error {
	if _, ok := users[u.Username]; ok {
		return ErrorUserAlreadyExists
	}
	users[u.Username] = u
	log.Println("User ", u.Username, "registered")
	return nil
}

func Get(username string) (User, error) {
	if user, ok := users[username]; ok {
		return user, nil
	}
	return User{}, ErrorUserNotFound
}

func Authenticate(username, password string) bool {
	user, err := Get(username)
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
