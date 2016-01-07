package main

import (
	"fmt"
	"github.com/asyoume/Auth/pkg/handler"
	"github.com/labstack/echo"
	"net/http"
)

var uhander = handler.UserHandler{}

func UserRegister(c *echo.Context) error {
	r, err := uhander.Register(c.Form("u"), "")
	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusInternalServerError, err.Error())
	}
}

func UserLogin(c *echo.Context) error {
	r, err := uhander.Register(c.Form("u"), "")

	fmt.Println(err)
	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusOK, "err")
	}
}
