package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"latih.in-be/config"
	"latih.in-be/controller"
	"latih.in-be/repository"
	"latih.in-be/route"
	"latih.in-be/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := gin.Default()

	route.UserRoutes(r, userController)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println("FATAL ERROR: Server failed to run:", err)
		os.Exit(1)
	}
}
