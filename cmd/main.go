package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/canermastan/teknofest2025-go-backend/internal/middleware"
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"github.com/canermastan/teknofest2025-go-backend/internal/config"
	"github.com/canermastan/teknofest2025-go-backend/internal/routes"
	"github.com/canermastan/teknofest2025-go-backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config yüklenemedi: %v", err)
	}
	// Connect to local database
	db, err := utils.ConnectDB(cfg.LocalDB)
	if err != nil {
		log.Fatalf("Local DB bağlantı hatası: %v", err)
	}
	log.Println("Veritabanına başarıyla bağlanıldı.")

	if err := db.AutoMigrate(
		&model.Review{},
		&model.Chest{},
		&model.CollectedChest{},
	); err != nil {
		log.Fatalf("Migrate işlemi başarısız: %v", err)
	}
	
	app.Use(middleware.LoggerMiddleware())
	routes.RegisterRoutes(app, db)
	
	// Server start and Gracefully shutdown server with a timeout
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		if err := app.Listen(":3001"); err != nil {
			log.Fatalf("Failed to start server: %v", err)	
		}
	}()
	
	receivedSignal := <-stop
	log.Printf("Received signal: %s, shutting down server...", receivedSignal)
	
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")
}