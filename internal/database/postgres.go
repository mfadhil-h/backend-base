package database

import (
	"backend-base/internal/model"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// ✅ Auto-migrate tables
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
