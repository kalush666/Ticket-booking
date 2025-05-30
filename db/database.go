package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kalush66/ticket-booking-project-v1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig,DBMigrator func(db*gorm.DB) error) *gorm.DB{
	uri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s port=5432",
		config.DBhost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode,
	)

	db,err := gorm.Open(postgres.Open(uri),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil{
		log.Fatalf("Unable to connect to the database: %e",err)
	}

	log.Info("Connected to db")

	if err:= DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate tables: %e",err)
	}

	return db
}