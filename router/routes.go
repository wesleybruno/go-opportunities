package router

import (
	docs "github.com/LagLabs/backend-go.git/docs"
	"github.com/LagLabs/backend-go.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {

	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)

	{
		v1.GET("/opennings", handler.ListAllOpenningHandler)
		v1.GET("/openning", handler.ListOneOpenningHandler)
		v1.POST("/openning", handler.CreateOpenningHandler)
		v1.PUT("/openning", handler.UpdateOpenningHandler)
		v1.DELETE("/openning", handler.DeleteOpenningHandler)
	}

	r.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
