package token

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/o5h/services/services/token"
	"github.com/o5h/services/services/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AccessResponse struct {
	AccessToken string `json:"access_token"`
}

// LoginHandler handles user login and returns a JWT token upon successful authentication
func LoginHandler(c echo.Context) error {
	var u user.User
	// Parse the request body to get username and password
	if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}

	// Authenticate the user
	if !user.Authenticate(u.Username, u.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}

	// Generate JWT token
	token, err := token.Create(u.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating token")
	}

	// Respond with the generated token
	return c.JSON(http.StatusOK, &AccessResponse{AccessToken: token})
}

// RefreshTokenHandler handles the refresh token request
func RefreshTokenHandler(c echo.Context) error {
	oldTokenString := c.Request().Header.Get("Authorization")

	// Check if the old token is provided
	if oldTokenString == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Old token not provided"})
	}

	// Extract the claims from the old token
	// Parse the token
	claims := token.AccessClaims{}
	_, err := jwt.ParseWithClaims(oldTokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return token.JWT_KEY, nil
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error extracting old token claims"})
	}

	// Check if the old token is expired
	if time.Now().After(claims.ExpiresAt.Time) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Old token has expired"})
	}

	// Create a new token for the same user
	newToken, err := token.Create(claims.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error creating new token"})
	}

	return c.JSON(http.StatusOK, &AccessResponse{AccessToken: newToken})
}

func RevokeHandler(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	token.Revoke(tokenString)
	return c.NoContent(http.StatusOK)
}
