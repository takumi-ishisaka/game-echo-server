package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Run() error {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!!!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.Logger.Fatal(e.Start(":1323"))
	return nil
}