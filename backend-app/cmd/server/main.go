package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/savindaJ/backend-app/internal/server"
	"go.uber.org/zap"

	_ "github.com/savindaJ/backend-app/docs" // Swagger docs
)

// @title           Go Backend API
// @version         1.0
// @description     A backend API server built with Go and Gin
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	// Setup logger BEFORE starting server
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("🚀 Starting server...")

	// This blocks - must be last!
	server.Start()
}
