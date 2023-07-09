package main

import (
	"fmt"
	"log"

	"github.com/takumi-ishisaka/game-echo-server/routes"
	"github.com/takumi-ishisaka/game-echo-server/infras"
)

func main() {
	db, _ := infras.DB.DB()
	defer db.Close()
	err := routes.Run()
	if err != nil {
		log.Fatal(fmt.Sprintf(`{"error":"%v"}`, err))
	}
}