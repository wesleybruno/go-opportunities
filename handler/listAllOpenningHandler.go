package handler

import (
	"net/http"

	"github.com/LagLabs/backend-go.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListAllOpenningHandler(ctx *gin.Context) {

	opennings := []schemas.Oppening{}

	if err := db.Find(&opennings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error to find all entity")
		return
	}

	sendSuccess(ctx, "ListAllOpenningHandler", opennings)

}
