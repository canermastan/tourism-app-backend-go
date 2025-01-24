package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"regexp"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// LoadTestConfig loads the database configuration for tests
func LoadTestConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnv("DB_LOCAL_HOST", "localhost"),
		Port:     getEnv("DB_LOCAL_PORT", "5432"),
		User:     getEnv("DB_LOCAL_USER", "testuser"),
		Password: getEnv("DB_LOCAL_PASSWORD", "testpassword"),
		Name:     getEnv("DB_LOCAL_NAME", "testdb"),
	}
}

func init() {
	projectName := regexp.MustCompile(`^(.*` + "teknofest2025-go-backend" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Println("No .env file found.")
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}

func ConnectTestDB(cfg *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func SetupTestDB() (*gorm.DB, error) {
	cfg := LoadTestConfig()
	db, err := ConnectTestDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to test database: %v", err)
	}
	return db, nil
}

func CleanupTestDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("Failed to get sqlDB from gorm.DB: %v", err)
	}
	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("Failed to close test database connection: %v", err)
	}
	return nil
}
