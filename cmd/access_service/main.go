package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/o5h/services/config"
	tokenHandlers "github.com/o5h/services/services/token/handlers"
	"github.com/o5h/services/services/token/middleware"
	userHandlers "github.com/o5h/services/services/user/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	e := echo.New()
	public := e.Group("")
	{
		public.POST("/user", userHandlers.RegisterHandler)
		public.POST("/access/login", tokenHandlers.LoginHandler)
	}

	protected := e.Group("")
	{
		protected.Use(middleware.ValidateTokenMiddleware)
		protected.GET("/user/details", userHandlers.DetailsHandler)
		protected.POST("/access/refresh", tokenHandlers.RefreshTokenHandler)
		protected.DELETE("/access/revoke", tokenHandlers.RevokeHandler)
	}

	if err := e.Start(fmt.Sprintf(":%d", conf.App.Port)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
