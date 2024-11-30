package main

import (
	"context"
	"fmt"
	"log"

	"github.com/canermastan/teknofest2025-go-backend/internal/config"
	"github.com/canermastan/teknofest2025-go-backend/internal/routes"
	"github.com/canermastan/teknofest2025-go-backend/internal/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config yüklenemedi: %v", err)
	}

	localDB, err := utils.ConnectDB(cfg.LocalDB)
	if err != nil {
		log.Fatalf("Local DB bağlantı hatası: %v", err)
	}
	defer localDB.Close()

	remoteDB, err := utils.ConnectDB(cfg.RemoteDB)
	if err != nil {
		log.Fatalf("Remote DB bağlantı hatası: %v", err)
	}
	defer remoteDB.Close()

	logger, err := utils.NewZapLogger(cfg, remoteDB)
	if err != nil {
		log.Fatalf("Logger başlatılamadı: %v", err)
	}

	// Örnek loglar
	ctx := context.Background()
	logger.Info(ctx, "Proje başlatıldı")
	logger.Error(ctx, "Bir hata oluştu")
	fmt.Println("Proje çalışıyor!")

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
