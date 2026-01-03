package main

import "github.com/yourname/go-backend/internal/server"
import "go.uber.org/zap"
import "github.com/joho/godotenv"
import "log"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	server.Start()
	logger, _ := zap.NewProduction()
	logger.Info("Server started")
	defer logger.Sync()
}
