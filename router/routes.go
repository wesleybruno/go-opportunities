package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")

	{
		v1.GET("/openning", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET openning",
			})
		})

		v1.GET("/opennings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET opennings",
			})
		})

		v1.POST("/openning", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "POST openning",
			})
		})

		v1.PUT("/openning", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "PUT openning",
			})
		})

		v1.DELETE("/openning", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DELETE openning",
			})
		})
	}

}
