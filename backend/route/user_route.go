package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
	"latih.in-be/middleware"
)

func UserRoutes(r *gin.Engine, user *controller.UserController) {
	users := r.Group("/user")
	{
		users.POST("/register", user.Register)
		users.POST("/login", user.Login)

		users.POST("/refresh", user.RefreshToken)

		usersAuth := users.Group("")
		usersAuth.Use(middleware.AuthMiddleware())
		{
			usersAuth.GET("/id/:id", user.GetById)
			usersAuth.GET("/email/:email", user.GetByEmail)
			usersAuth.GET("/nim/:nim", user.GetByNim)
			usersAuth.GET("/name/:name", user.GetByName)
			usersAuth.GET("/role/:role", user.GetByRole)
			usersAuth.GET("/", user.GetAll)
			usersAuth.PUT("/:id", user.Update)
			usersAuth.PUT("/password/:id", user.ChangePassword)
			usersAuth.DELETE("/id/:id", user.Delete)
			usersAuth.PUT("/role/:id", user.ChangeRole)
		}
	}
}
