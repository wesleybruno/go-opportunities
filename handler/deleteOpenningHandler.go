package handler

import (
	"net/http"

	"github.com/LagLabs/backend-go.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpenningHandler(ctx *gin.Context) {
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

	if err := db.Delete(&openning, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error to delete entity not found")
		return
	}

	sendSuccess(ctx, "DeleteOpenningHandler", openning)

}
