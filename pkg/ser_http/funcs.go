package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func UserRegister(c *echo.Context) error {
	r, err := uhander.Register(c.Query("u"), "")
	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusOK, "err")
	}
}

func UserLogin(c *echo.Context) error {
	r, err := uhander.Register(c.Form("name"), "")

	fmt.Println(err)
	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusOK, "err")
	}
}
