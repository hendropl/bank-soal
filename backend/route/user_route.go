package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
	"latih.in-be/middleware"
)

func UserRoutes(r *gin.Engine, user *controller.UserController) {
	users := r.Group("/users")
	{
		users.POST("/register", user.Register)

		usersAuth := users.Group("")
		usersAuth.Use(middleware.AuthMiddleware())
		{
			usersAuth.GET("/:id")
		}
	}
}
