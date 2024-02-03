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

	e.POST("/user", users.RegisterHandler)
	group := e.Group("/user")
	{
		group.Use(access.ValidateTokenMiddleware)
		group.GET("/details", users.DetailsHandler)
	}

	e.POST("/access/login", access.LoginHandler)
	refresh := e.Group("/access")
	{
		refresh.Use(access.ValidateTokenMiddleware)
		refresh.POST("/refresh", access.RefreshTokenHandler)
		refresh.DELETE("/revoke", access.InvalidateTokenHandler)

	}

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
