package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusabreusouza/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	logger.Infof("Received request to create opening: %+v", request)

	request.Validate()
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation failed: %v", err)
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Title:    request.Title,
		Company:  request.Company,
		Location: request.Location,
		Salary:   request.Salary,
		Remote:   *request.Remote,
		Role:     request.Role,
		Link:     request.Link,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Failed to create opening: %v", err)
		SendError(ctx, http.StatusInternalServerError, "Failed to create opening")
		return
	}

	SendSuccess(ctx, "Create opening", http.StatusCreated, opening)
}
