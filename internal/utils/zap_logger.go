package utils

import (
	"context"
	"database/sql"

	//"your_project/internal/model"

	"github.com/canermastan/teknofest2025-go-backend/internal/config"
	"go.uber.org/zap"
)

type ZapLogger struct {
	zapLogger *zap.Logger
	remoteDB  *sql.DB
	devName   string
}

func NewZapLogger(cfg *config.Config, remoteDB *sql.DB) (*ZapLogger, error) {
	zapLogger := zap.NewExample().With(zap.String("developer", cfg.DevName))
	return &ZapLogger{
		zapLogger: zapLogger,
		remoteDB:  remoteDB,
		devName:   cfg.DevName,
	}, nil
}

func (l *ZapLogger) Info(ctx context.Context, message string) {
	l.zapLogger.Info(message)
	/*log := model.Log{
		DeveloperName: l.devName,
		LogMessage:    message,
		LogLevel:      "INFO",
		CreatedAt:     time.Now(),
	}
	log.Insert(ctx, l.remoteDB, boil.Infer())*/
}

func (l *ZapLogger) Error(ctx context.Context, message string) {
	l.zapLogger.Error(message)
	/*log := model.Log{
		DeveloperName: l.devName,
		LogMessage:    message,
		LogLevel:      "ERROR",
		CreatedAt:     time.Now(),
	}
	log.Insert(ctx, l.remoteDB, boil.Infer())*/
}
