package utils

import (
	"database/sql"

	"go.uber.org/zap"
)

type ZapLogger struct {
	zapLogger *zap.Logger
	remoteDB  *sql.DB
	devName   string
}

// TODO: develop this
/*
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
	log := model.Log{
		DeveloperName: l.devName,
		LogMessage:    message,
		LogLevel:      "INFO",
		CreatedAt:     time.Now(),
	}
	log.Insert(ctx, l.remoteDB, boil.Infer())
}

func (l *ZapLogger) Error(ctx context.Context, message string) {
	l.zapLogger.Error(message)
	log := model.Log{
		DeveloperName: l.devName,
		LogMessage:    message,
		LogLevel:      "ERROR",
		CreatedAt:     time.Now(),
	}
	log.Insert(ctx, l.remoteDB, boil.Infer())
}*/
