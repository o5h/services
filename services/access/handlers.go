package access

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/o5h/services/services/users"
)

// LoginHandler handles user login and returns a JWT token upon successful authentication
func LoginHandler(c echo.Context) error {
	var user users.User
	// Parse the request body to get username and password
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}

	// Authenticate the user
	if !GetService().AuthenticateUser(user.Username, user.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}

	// Generate JWT token
	token, err := GetService().CreateToken(user.Username, time.Minute*5)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating token")
	}

	// Respond with the generated token
	_, err = c.Response().Write([]byte(token))
	return err
}
