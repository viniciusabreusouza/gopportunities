package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/viniciusabreusouza/gopportunities/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")

	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Error("Failed to find opening", "error", err)
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	if request.Title != "" {
		opening.Title = request.Title
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary.GreaterThanOrEqual(decimal.Zero) {
		opening.Salary = request.Salary
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if err := db.Save(&opening).Error; err != nil {
		logger.Error("Failed to update opening", "error", err)
		SendError(ctx, http.StatusInternalServerError, "Failed to update opening")
		return
	}
	logger.Info("Opening updated successfully", "id", opening.ID)
	SendSuccess(ctx, "Update opening", http.StatusOK, opening)
}
