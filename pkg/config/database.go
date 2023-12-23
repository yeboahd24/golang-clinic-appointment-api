package config

import (
    "clinic-appointment-system/pkg/model"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file", err)
    }

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSLMODE"),
    )
    fmt.Println("DSN:", dsn)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    // AutoMigrate for other models as needed
    db.AutoMigrate(&model.User{}, &model.Appointment{})

    return db
}

func CloseDatabaseConnection(db *gorm.DB) {
    sqlDB, err := db.DB()
    if err != nil {
        panic("Failed to close database connection: " + err.Error())
    }
    sqlDB.Close()
}

