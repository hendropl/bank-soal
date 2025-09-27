package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
)

func UserRoutes(r *gin.Engine, user *controller.UserController) {
	// api := r.Group("/api")
	users := r.Group("/users")
	{
		users.POST("/register", user.Register)
	}
}
