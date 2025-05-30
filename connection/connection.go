package connection

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/alwialdi9/be_auth-jajanskuy/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewConnection() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Error("Database connection returned an error!")
	}

	if err = DB.AutoMigrate(&models.User{}); err != nil {
		log.WithError(err).Error("Failed to auto-migrate User model!")
	}
	log.Info("Postgres Database connection established successfully! and success migrate 🚀🚀")

	// If there is an error, return the error
	// return fmt.Errorf("failed to connect: %v", err)
}

func CheckConnection() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = DB.WithContext(ctx).Exec("SELECT 1").Error
	if err != nil {
		log.Fatalf("Postgres Database failed: %v", err)
	} else {
		fmt.Println("✅ Postgres Database connection OK! 🚀🚀")
	}
}
