package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
)

func ExamRoutes(r *gin.Engine, exam *controller.ExamController) {
	exams := r.Group("/exam")
	{
		exams.POST("/", exam.Create)
		exams.GET("/", exam.GetAll)
		exams.GET("/id/:id", exam.GetById)
		exams.PUT("/:id", exam.Update)
		exams.DELETE("/:id", exam.Delete)
	}
}
