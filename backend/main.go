package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"latih.in-be/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	config.InitDB()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println("FATAL ERROR: Server failed to run:", err)
		os.Exit(1)
	}
}
