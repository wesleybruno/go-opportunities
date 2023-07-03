package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST openning",
	})
}

func ListOneOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST openning",
	})
}

func ListAllOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST openning",
	})
}

func UpdateOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST openning",
	})
}

func DeleteOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST openning",
	})
}
