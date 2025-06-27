package handler

import "github.com/gin-gonic/gin"

func DeleteOpeningHandler(ctx *gin.Context) {
	request := struct {
		Title string `json:"title" binding:"required"`
	}{}
	ctx.BindJSON(&request)

	logger.Info("Received request to create opening", "title", request.Title)
}
