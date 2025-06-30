package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusabreusouza/gopportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Error("Failed to list openings", "error", err)
		SendError(ctx, http.StatusInternalServerError, "Failed to list openings")
		return
	}
	if len(openings) == 0 {
		SendSuccess(ctx, "No openings found", http.StatusOK, nil)
		return
	}
	SendSuccess(ctx, "List openings", http.StatusOK, openings)
}
