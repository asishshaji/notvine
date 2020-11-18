package controller

import (
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/asishshaji/notvine/app/entity"
	"github.com/asishshaji/notvine/app/usecase"
	"github.com/asishshaji/notvine/app/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AppController struct {
	appusecase usecase.AppUsecase
	bucket     *storage.BucketHandle
}

func NewAppController(usecase usecase.AppUsecase, bucket *storage.BucketHandle) ControllerInterface {
	return &AppController{
		appusecase: usecase,
		bucket:     bucket,
	}

}

// Signup creates user in the database
func (a *AppController) Signup(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	_, err := a.appusecase.Signup(c.Request().Context(), username, password)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil

}

// Login sends token to user
func (a *AppController) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := a.appusecase.Login(c.Request().Context(), username, password)

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

// CreatePost creates new post
func (a *AppController) CreatePost(c echo.Context) error {

	file, err := c.FormFile("video_file")
	caption := c.FormValue("caption")

	userJWT := c.Get("user").(*jwt.Token)
	claims := userJWT.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]error{
			"error": err,
		})
	}
	link, err1 := utils.UploadVideo(file, a.bucket)

	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err1.Error(),
		})
	}

	user, err3 := a.appusecase.GetUser(c.Request().Context(), username)

	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err3.Error(),
		})
	}

	post := entity.Post{
		URL:        link,
		Owner:      user,
		Caption:    caption,
		LikesCount: 0,
		CreatedAt:  time.Now(),
	}

	postErr := a.appusecase.CreatePost(c.Request().Context(), &post)

	if postErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": postErr.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Created Post",
	})
}
