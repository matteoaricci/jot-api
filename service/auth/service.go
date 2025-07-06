package auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func WriteAuthCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func ReadAuthCookie(c echo.Context) error {
	cookie, err := c.Cookie("auth")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}
