package router

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "List of users efae ",
			})
		})
		// Add more user-related routes here
	}
}
