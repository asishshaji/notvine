package controller

import (
	"net/http"
	"time"

	"github.com/asishshaji/notvine/app/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AppController struct {
	appusecase usecase.AppUsecase
}

func NewAppController(usecase usecase.AppUsecase) *AppController {
	return &AppController{
		appusecase: usecase,
	}
}

func (a *AppController) Signup(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := a.appusecase.Signup(c.Request().Context(), username, password)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 240).Unix()

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}

func (a *AppController) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := a.appusecase.Login(c.Request().Context(), username, password)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// func (a *AppController) CreatePost(c echo.Context) (entity.Post, error) {}
// func (a *AppController) GetUser(c echo.Context) (entity.User, error)    {}
