package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/takumi-ishisaka/game-echo-server/infras"
)

type User struct {
	Name string `json:"name"`
}

func UserRegister(c echo.Context) error {
	db, _ := infras.DB.DB()
	defer db.Close()

	err := db.Ping()

	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	} else {
		return c.String(http.StatusOK, "DB接続しました")
	}
}