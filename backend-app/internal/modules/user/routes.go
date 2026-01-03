package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all user routes
func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	// Initialize dependencies
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	// User routes
	users := router.Group("/users")
	{
		// Public routes
		users.POST("/register", handler.Register)
		users.POST("/login", handler.Login)

		// Protected routes (add auth middleware later)
		users.GET("", handler.GetAll)
		users.GET("/:id", handler.GetByID)
		users.PUT("/:id", handler.Update)
		users.DELETE("/:id", handler.Delete)
	}
}
