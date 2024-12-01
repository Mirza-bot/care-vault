package config

import (
	"fmt"
	"log"
	"care-vault/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Load() {
    viper.AutomaticEnv()
    // viper.SetDefault("APP_PORT", "8080")
}

func ConnectDatabase() {
    var err error

    dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
        viper.GetString("POSTGRES_USER"),
        viper.GetString("POSTGRES_PASSWORD"),
        viper.GetString("POSTGRES_DB"),
    )

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Unable to connect to database %v \n", err)
    }

    fmt.Println("Connected to PostgreSQL!")

    Migrate(DB)
}



func Migrate(db *gorm.DB) {
    models := []interface{}{
        &models.User{},
    }

    for _, model := range models {
        if err := db.AutoMigrate(model); err != nil {
            log.Printf("Failed to migrate model %T: %v", model, err)
        } else {
            log.Printf("Successfully migrated model %T", model)
        }
    }
}


