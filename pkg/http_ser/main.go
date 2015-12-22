package main

import (
	"github.com/asyoume/Auth/pkg/handler"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"net/http"
)

var uhander = handler.UserHandler{}

func UserRegister(c *echo.Context) error {
	r, err := uhander.Register(c.Query("u"), "")
	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusOK, "err")
	}
}

func UserLogin(c *echo.Context) error {
	r, err := uhander.Register(c.Value("u").(string), "")
	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusOK, "err")
	}
}

func main() {
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Use(cors.Default().Handler)

	// Routes
	e.Get("/Register", UserRegister)
	e.Get("/Register", UserRegister)
	e.Run(":9091")
}
