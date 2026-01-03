package database

import (
	"fmt"
	"log"

	"github.com/savindaJ/backend-app/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect initializes the database connection with auto database creation
func Connect(cfg *config.Config) *gorm.DB {
	// First, connect without database name to create it if not exists
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
	)

	// Connect to MySQL server (without database)
	tempDB, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to MySQL server: %v", err)
	}

	// Create database if not exists
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", cfg.DBName)
	if err := tempDB.Exec(createDBSQL).Error; err != nil {
		log.Fatalf("❌ Failed to create database: %v", err)
	}
	log.Printf("✅ Database '%s' is ready", cfg.DBName)

	// Close temp connection
	sqlDB, _ := tempDB.DB()
	sqlDB.Close()

	// Now connect to the actual database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Configure logger based on environment
	logLevel := logger.Info
	if cfg.AppEnv == "production" {
		logLevel = logger.Error
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	log.Println("✅ Database connected successfully")
	return DB
}

// AutoMigrate runs auto migration for all models
func AutoMigrate(db *gorm.DB, models ...interface{}) {
	log.Println("🔄 Running auto migrations...")
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}
	log.Println("✅ Migrations completed successfully")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
