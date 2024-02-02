package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/o5h/services/services/access"
	"github.com/o5h/services/services/users"
)

func main() {
	e := echo.New()

	e.POST("/user", users.SigninHandler)
	e.POST("/access", access.LoginHandler)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
