package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/canermastan/teknofest2025-go-backend/internal/config"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
