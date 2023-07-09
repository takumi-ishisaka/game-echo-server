package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/takumi-ishisaka/game-echo-server/infras"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func UserRegister(c echo.Context) error {
	// if err := c.Bind(&user); err != nil {
	// 	return err
	// }
	user := &User{
		Id: "test",
		Name: "test",
	}

	result := infras.DB.Create(user)

	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "ユーザー作成に失敗しました")
	} else {
		return c.String(http.StatusOK, "ユーザーを作成しました")
	}
}