package database

import (
	"log"

	"github.com/Mokarama/assignment-6-spotsync-api/config"
	"github.com/Mokarama/assignment-6-spotsync-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := config.GetEnv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// Auto migrate database tables
	err = DB.AutoMigrate(
		&models.User{},
		&models.ParkingZone{},
		&models.Reservation{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("✅ Database connected successfully")
	log.Println("✅ Database migration completed successfully")
}
