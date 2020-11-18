package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerInterface interface {
	Login(c echo.Context) error
	Signup(c echo.Context) error

	CreatePost(c echo.Context) error
}
