package database

import (
	"backend-base/internal/model"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

// DSN builds the Postgres connection string
func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode,
	)
}

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
	if err := migrate(db); err != nil {
		log.Fatal().Err(err).Msg("❌ Database migration failed")
		return nil, err
	}

	// ✅ Run seeding right after migration
	if err := seed(db); err != nil {
		log.Fatal().Err(err).Msg("❌ Database seeding failed")
	}

	log.Info().Msg("✅ Database connected and migrated successfully.")
	return db, nil
}

func ConnectPostgres(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// ✅ Auto-migrate tables
	if err := migrate(db); err != nil {
		log.Fatal().Err(err).Msg("❌ Database migration failed")
		return nil, err
	}

	// ✅ Run seeding right after migration
	if err := seed(db); err != nil {
		log.Fatal().Err(err).Msg("❌ Database seeding failed")
	}

	log.Info().Msg("✅ Database connected and migrated successfully.")
	return db, nil
}

// migrate runs all model migrations automatically
func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		// You can append more models here as you expand
	)
}

func seed(db *gorm.DB) error {
	log.Info().Msg("🌱 Starting database seeding...")

	// Example: Seed an admin user if not exists
	adminEmail := "admin@example.com"
	var count int64
	db.Model(&model.User{}).Where("email = ?", adminEmail).Count(&count)

	if count == 0 {
		admin := model.User{
			Name:     "Administrator",
			Email:    adminEmail,
			Password: "$2a$10$Bn1hQEY53FUXe3LjWWcveuwK6TIBWNPpn8U.L5kKq5nM7W7pQbgHS", // bcrypt("admin123")
		}
		if err := db.Create(&admin).Error; err != nil {
			return err
		}
		log.Info().Msg("👑 Admin user seeded: admin@example.com / admin123")
	} else {
		log.Info().Msg("👑 Admin user already exists, skipping seed.")
	}

	log.Info().Msg("🌱 Database seeding completed successfully.")
	return nil
}
