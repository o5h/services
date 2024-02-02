package users

import (
	"errors"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type API interface {
	CreateUser(user User) error
	GetUser(username string) (User, error)
}

var (
	ErrorUserAlreadyExists = errors.New("User already exists")
	ErrorUserNotFound      = errors.New("User not found")
)

var instance = &inMemory{
	users: make(map[string]User)}

func GetService() API {
	return instance
}

// In a real-world scenario, you would store the hashed password in a secure data store
type inMemory struct {
	users map[string]User
}

func (serv *inMemory) CreateUser(user User) error {
	if _, ok := serv.users[user.Username]; ok {
		return ErrorUserAlreadyExists
	}
	serv.users[user.Username] = user
	log.Println("User ", user.Username, "registered")
	return nil
}

func (storage *inMemory) GetUser(username string) (User, error) {
	if user, ok := storage.users[username]; ok {
		return user, nil
	}
	return User{}, ErrorUserNotFound
}
