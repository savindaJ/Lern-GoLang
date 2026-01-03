package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yourname/go-backend/internal/config"
)

func Start() {
	cfg := config.Load()

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
