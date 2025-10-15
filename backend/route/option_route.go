package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
)

func OptionRoutes(r *gin.Engine, option *controller.OptionController) {
	options := r.Group("/option")
	{
		options.POST("/", option.Create)
		options.GET("/", option.GetAll)
		options.GET("/id/:id", option.GetById)
		options.PUT("/:id", option.Update)
		options.DELETE("/:id", option.Delete)
	}
}
