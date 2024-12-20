package main

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/middleware"
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"log"

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
	); err != nil {
		log.Fatalf("Migrate işlemi başarısız: %v", err)
	}
	// Connect to remote database
	/*_, err = utils.ConnectDB(cfg.RemoteDB)
	if err != nil {
		log.Fatalf("Remote DB bağlantı hatası: %v", err)
	}*/

	/*logger, err := utils.NewZapLogger(cfg, remoteDB)
	if err != nil {
		log.Fatalf("Logger başlatılamadı: %v", err)
	}

	// Örnek loglar
	ctx := context.Background()
	logger.Info(ctx, "Proje başlatıldı")
	logger.Error(ctx, "Bir hata oluştu")*/

	app.Use(middleware.LoggerMiddleware())
	routes.RegisterRoutes(app, db)

	app.Listen(":3001")
}
