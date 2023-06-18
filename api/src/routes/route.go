package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/takumi-ishisaka/game-echo-server/controllers"
)

func Run() error {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!!!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	// user
	e.GET("/user/register", controllers.UserRegister)

	e.Logger.Fatal(e.Start(":1323"))
	return nil
}