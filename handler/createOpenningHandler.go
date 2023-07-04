package handler

import (
	"net/http"

	"github.com/LagLabs/backend-go.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpenningHandler(ctx *gin.Context) {

	request := CreateOpenningRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request received: %v", request)

	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	openning := schemas.Oppening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&openning).Error; err != nil {
		logger.Errorf("Error creating openning: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error creating oppnening on db")
		return
	}

	sendSuccess(ctx, "CreateOpenningHandler", openning)

}
