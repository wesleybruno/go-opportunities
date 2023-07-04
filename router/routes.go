package router

import (
	"github.com/LagLabs/backend-go.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {

	handler.InitializeHandler()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/opennings", handler.ListAllOpenningHandler)
		v1.GET("/openning", handler.ListOneOpenningHandler)
		v1.POST("/openning", handler.CreateOpenningHandler)
		v1.PUT("/openning", handler.UpdateOpenningHandler)
		v1.DELETE("/openning", handler.DeleteOpenningHandler)
	}

}
