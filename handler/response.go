package handler

import (
	"fmt"
	"net/http"

	"github.com/LagLabs/backend-go.git/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation success: %v ", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}
type ShowOppeningResponse struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                     `json:"message"`
	Data    []schemas.OppeningResponse `json:"data"`
}
type UpdateOppeningResponse struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}
