package handler

import (
	"net/http"

	"github.com/LagLabs/backend-go.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListAllOpenningHandler(ctx *gin.Context) {

	opennings := []schemas.Oppening{}

	if err := db.Find(&opennings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error to find all entity")
		return
	}

	sendSuccess(ctx, "ListAllOpenningHandler", opennings)

}
