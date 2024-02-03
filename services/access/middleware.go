package access

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ValidateTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		// Check if the token is provided
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
		}

		if !GetService().IsTokenValid(tokenString) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Token was revoked."})
		}

		// Parse the token
		claims := &AccessTokenClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
			}
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Bad request"})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Token is not valid"})
		}

		// Set the user from the token claims into the Echo context
		c.Set("username", claims.Username)

		return next(c)
	}
}
