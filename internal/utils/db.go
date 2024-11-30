package utils

import (
	"database/sql"
	"fmt"

	"github.com/canermastan/teknofest2025-go-backend/internal/config"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	return sql.Open("postgres", dsn)
}
