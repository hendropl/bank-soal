package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
)

func QuestionRoutes(r *gin.Engine, question *controller.QuestionController) {
	questions := r.Group("/question")
	{
		questions.POST("/", question.Create)
		questions.GET("/", question.GetAll)
		questions.GET("/id/:id", question.GetById)
		questions.PUT("/:id", question.Update)
		questions.DELETE("/:id/:userId", question.Delete)
		questions.POST("/options", question.CreateWithOptions)
		questions.POST("/json", question.CreateFromJson)
	}
}
