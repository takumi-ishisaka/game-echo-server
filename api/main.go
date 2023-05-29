package main

import (
	"fmt"
	"log"

	"github.com/takumi-ishisaka/game-echo-server/routes"
)

func main() {
	err := routes.Run()
	if err != nil {
		log.Fatal(fmt.Sprintf(`{"error":"%v"}`, err))
	}
}