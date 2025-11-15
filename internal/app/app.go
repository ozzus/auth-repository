package app

import (
	"log/slog"
	"time"

	grpcapp "ozzus/auth-repository/internal/app/grpc"
	"ozzus/auth-repository/internal/domain"
	"ozzus/auth-repository/internal/services/auth"
	"ozzus/auth-repository/internal/storage/psql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, dsn string, tokenTTL time.Duration, appSecret string) *App {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.User{})

	userRepo := psql.NewUserRepository(db)

	authService := auth.New(log, userRepo, tokenTTL, appSecret)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
