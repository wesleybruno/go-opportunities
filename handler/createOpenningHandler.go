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
