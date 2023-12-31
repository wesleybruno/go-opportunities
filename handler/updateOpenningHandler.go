package handler

import (
	"net/http"

	"github.com/LagLabs/backend-go.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Param opening body UpdateOpenningRequest true "Opening data to Update"
// @Success 200 {object} UpdateOppeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpenningHandler(ctx *gin.Context) {
	//Find oppnening
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

	//Update new oppening data
	update := UpdateOpenningRequest{}

	ctx.BindJSON(&update)

	if err := update.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Update opening
	if update.Role != "" {
		openning.Role = update.Role
	}

	if update.Company != "" {
		openning.Company = update.Company
	}

	if update.Location != "" {
		openning.Location = update.Location
	}

	if update.Remote != nil {
		openning.Remote = *update.Remote
	}

	if update.Link != "" {
		openning.Link = update.Link
	}

	if update.Salary > 0 {
		openning.Salary = update.Salary
	}

	// Update Opennig
	if err := db.Save(&openning).Error; err != nil {
		logger.Errorf("validation error: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "UpdateOpenningHandler", http.StatusOK, openning)
}
