package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/o5h/services/config"
	"github.com/o5h/services/services/access"
	"github.com/o5h/services/services/users"
	log "github.com/sirupsen/logrus"
)

func main() {

	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
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

	if err := e.Start(fmt.Sprintf(":%d", conf.App.Port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
