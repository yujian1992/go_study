package main

import (
	"os"

	"go_study/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewService()
	server.Run(":" + port)
}
