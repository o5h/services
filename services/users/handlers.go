package users

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SigninHandler(c echo.Context) error {
	var user User

	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}
	// Generate hashed password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error hashing password")
	}
	user.Password = base64.StdEncoding.EncodeToString([]byte(hashedPassword))
	GetService().CreateUser(user)
	return nil
}
