package app

import (
	"net/http"

	"github.com/asishshaji/notvine/app/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	e    *echo.Echo
	port string
}

// NewApp creates new app
func NewApp(port string, controller controller.AppController) *App {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("randomstring"), //read from .env file
	// }))

	e.POST("/signup", controller.Signup)

	return &App{
		e:    e,
		port: port,
	}

}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// RunServer starts the server
func (a *App) RunServer() {
	a.e.Logger.Fatal(a.e.Start(a.port))

}
