package router

import (
	"net/http"

	models "example.com/gin-first-project/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func SetupAuthRoutes(r *gin.Engine, db *gorm.DB) {
	authGroup := r.Group("/auth")
	{
		// register user route
		authGroup.POST("/register", func(c *gin.Context) {
			var userJson RegisterUser
			if err := c.BindJSON(&userJson); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			var user models.User
			user.Email = userJson.Email
			user.Password = userJson.Password
			err := models.CreateUser(db, &user) // pass pointer of data to Create
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			}
			// Save the user to the database
			c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
		})
		// login user route
		authGroup.POST("/login", func(c *gin.Context) {
			var userJson LoginUser
			if err := c.BindJSON(&userJson); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			var user models.User
			err := models.GetUserByEmail(db, userJson.Email, &user) // pass pointer of data to Create
			if err != nil && len(user.Email) == 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			// Save the user to the database
			c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
		})
	}
}
