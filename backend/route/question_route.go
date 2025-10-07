package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
)

func QuestionRoutes(r *gin.Engine, question *controller.QuestionController) {
	questions := r.Group("/exam")
	{
		questions.POST("/", question.Create)
		questions.GET("/", question.GetAll)
		questions.GET("/id/:id", question.GetById)
		questions.PUT("/", question.Update)
		questions.DELETE("/:id", question.Delete)
	}
}
