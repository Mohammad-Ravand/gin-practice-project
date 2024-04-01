package router
import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	// Setup routes from different files
	SetupHomeRoutes(r)
	SetupUserRoutes(r)
	SetupAuthRoutes(r, db)
	// Add more Setup functions for other route files
	return r
}
