package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := InitFirebase()
	if app == nil {
		log.Fatal("Error initializing Firebase")
	}

	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
