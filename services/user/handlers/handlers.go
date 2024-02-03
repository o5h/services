package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/o5h/services/services/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c echo.Context) error {
	var u user.User

	if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}
	// Generate hashed password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error hashing password")
	}
	u.Password = base64.StdEncoding.EncodeToString([]byte(hashedPassword))
	user.Create(u)
	return nil
}

func DetailsHandler(c echo.Context) error {
	username := c.Get("username").(string)
	user, err := user.Get(username)
	if err != nil {
		return err
	}
	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
