package config

import (
	"log"
	"uplink-go/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(cfg *Config) *gorm.DB {
	var logLevel logger.LogLevel
	logLevel = logger.Info

	db, error := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if error != nil {
		log.Fatal("Failed to connect to database:", error)
	}

	log.Println("✅ Database connected")
	return db
}

func AutoMigrate(db *gorm.DB) {
	error := db.AutoMigrate(
		&domain.User{},
	)
	
	if error != nil {
		log.Fatal("Failed to migrate database:", error)
	}
	
	log.Println("✅ Database migrated")
}