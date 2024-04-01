package router

import (
	"github.com/gin-gonic/gin"
)

func SetupHomeRoutes(r *gin.Engine) {
	userGroup := r.Group("/")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "home page",
			})
		})
		// Add more user-related routes here
	}
}
