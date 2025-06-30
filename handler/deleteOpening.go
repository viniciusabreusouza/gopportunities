package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusabreusouza/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First("id = ?", id).Find(&opening).Error; err != nil {
		logger.Error("Failed to find opening", "error", err)
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Error("Failed to delete opening", "error", err)
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Failed to delete opening with id %s", id))
		return
	}

	SendSuccess(ctx, "Delete opening", http.StatusOK, nil)
}
