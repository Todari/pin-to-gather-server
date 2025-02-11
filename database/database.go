package database

import (
	"fmt"
	"log"

	"github.com/Todari/pin-to-gather-server/config"
	"github.com/Todari/pin-to-gather-server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    config.LoadConfig()

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        config.AppConfig.DatabaseHost,
        config.AppConfig.DatabaseUser,
        config.AppConfig.DatabasePassword,
        config.AppConfig.DatabaseName,
        config.AppConfig.DatabasePort,
        config.AppConfig.SSLMode,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Failed to connect to the database:", err)
    }

    log.Println("✅ Database connected successfully")
    db.AutoMigrate(&models.Board{})
    DB = db
}