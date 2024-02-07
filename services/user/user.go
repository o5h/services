package user

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/o5h/services/db"
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

func Create(u User) error {
	tx := db.BeginTx()
	defer db.Commit(tx)

	id := db.UserInsert(tx, &db.User{
		UserName:     u.Username,
		PasswordHash: u.Password})

	log.Println("User ", u.Username, id, "registered")
	return nil
}

func Get(username string) (User, error) {
	tx := db.BeginReadOnlyTx()
	defer db.Commit(tx)

	u := db.UserGetByUserName(tx, username)
	if u == nil {
		return User{}, ErrorUserNotFound
	}
	return User{
			Username: u.UserName,
			Password: u.PasswordHash},
		nil
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
