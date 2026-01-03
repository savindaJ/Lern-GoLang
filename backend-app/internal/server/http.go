package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/savindaJ/backend-app/internal/config"
	"github.com/savindaJ/backend-app/internal/database"
	"github.com/savindaJ/backend-app/internal/modules/user"
)

func Start() {
	cfg := config.Load()

	// Connect to database
	db := database.Connect(cfg)

	// Auto migrate models
	database.AutoMigrate(db, &user.User{})

	// Set Gin mode
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Health check endpoint
	// @Summary      Health check
	// @Description  Check if the server is running
	// @Tags         health
	// @Produce      json
	// @Success      200  {object}  map[string]string
	// @Router       /health [get]
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Register module routes
		user.RegisterRoutes(v1, db)
	}

	log.Printf("🚀 Server starting on port %s", cfg.AppPort)
	log.Printf("📚 Swagger docs: http://localhost:%s/swagger/index.html", cfg.AppPort)

	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
