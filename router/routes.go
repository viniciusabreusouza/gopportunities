package router

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusabreusouza/gopportunities/handler"
)

func initializeRouter(router *gin.Engine) {
	handler.InitializeHandler()
	// Initialize the logger and database connection
	r := router.Group("/api/v1")
	{
		// r.GET("/health", handler.CreateHealthHandler)
		r.POST("/opening", handler.CreateOpeningHandler)
		// r.PUT("/opening", handler.UpdateOpeningHandler)
		r.DELETE("/opening", handler.DeleteOpeningHandler)
		// r.GET("/openings", handler.ListOpeningHandler)
	}
}
