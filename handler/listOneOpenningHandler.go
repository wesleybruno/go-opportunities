package handler

import (
	"net/http"

	"github.com/LagLabs/backend-go.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOppeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ListOneOpenningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	openning := schemas.Oppening{}

	if err := db.First(&openning, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Entity not found")
		return
	}

	sendSuccess(ctx, "ListOneOpenningHandler", http.StatusOK, openning)
}
