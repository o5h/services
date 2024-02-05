package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/o5h/services/config"
	tokenHandlers "github.com/o5h/services/services/token/handlers"
	tokenMiddlewarr "github.com/o5h/services/services/token/middleware"
	userHandlers "github.com/o5h/services/services/user/handlers"
)

var (
	server *echo.Echo
)

func Start(ctx context.Context, cancel context.CancelFunc) {

	go func() {
		<-ctx.Done()
		Shutdown()
	}()

	cfg := ctx.Value(config.ContextKey).(*config.Config)
	server = echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	r := server.Group("/internal")
	{
		r.GET("/shutdown", func(c echo.Context) error {
			defer cancel()
			return c.String(http.StatusOK, "Shutdown")
		})

		r.GET("/config", func(c echo.Context) error { return c.JSON(http.StatusOK, cfg) })

	}

	public := server.Group("")
	{
		public.POST("/user", userHandlers.RegisterHandler)
		public.POST("/access/login", tokenHandlers.LoginHandler)
	}

	protected := server.Group("")
	{
		protected.Use(tokenMiddlewarr.ValidateTokenMiddleware)
		protected.GET("/user/details", userHandlers.DetailsHandler)
		protected.POST("/access/refresh", tokenHandlers.RefreshTokenHandler)
		protected.DELETE("/access/revoke", tokenHandlers.RevokeHandler)
	}

	server.Logger.Info("Start")
	if err := server.Start(fmt.Sprintf("%v", cfg.App.Address)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func Shutdown() {
	server.Shutdown(context.Background())
}
