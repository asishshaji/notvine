package app

import (
	"github.com/asishshaji/notvine/app/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App creates the starting point of the server
type App struct {
	e    *echo.Echo
	port string
}

// NewApp creates new app
func NewApp(port string, controller controller.AppController) *App {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)

	return &App{
		e:    e,
		port: port,
	}

}

// RunServer starts the server
func (a *App) RunServer() {
	a.e.Logger.Fatal(a.e.Start(a.port))

}
