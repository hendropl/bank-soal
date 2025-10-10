package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
)

func ScoreRoutes(r *gin.Engine, score *controller.ScoreController) {
	scores := r.Group("/score")
	{
		scores.POST("/", score.Create)
		scores.GET("/all/:qId", score.GetAll)
		scores.GET("/id/:id", score.GetById)
		scores.PUT("/:id", score.Update)
		scores.DELETE("/:id", score.Delete)
	}
}
