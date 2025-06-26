package router

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusabreusouza/gopportunities/handler"
)

func initializeRouter(router *gin.Engine) {
	r := router.Group("/api/v1")
	{
		r.GET("/health", handler.CreateHealthHandler)
		r.POST("/opening", handler.CreateOpeningHandler)
		r.PUT("/opening", handler.UpdateOpeningHandler)
		r.DELETE("/opening", handler.UpdateOpeningHandler)
		r.GET("/openings", handler.ListOpeningHandler)
	}
}
