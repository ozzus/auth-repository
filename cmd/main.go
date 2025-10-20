package main

import (
	"fmt"
	"log/slog"
	"os"

	"ozzus/auth-repository/internal/app"
	"ozzus/auth-repository/internal/config"
	"ozzus/auth-repository/internal/lib/logger/slogpretty"

	"github.com/joho/godotenv"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting application")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("postgresql://postgres:%s@db.qdsaggkokumsosfwekae.supabase.co:5432/postgres", os.Getenv("DB_PASS"))
	application := app.New(log, cfg.GRPC.Port, dsn, cfg.TokenTTL, os.Getenv("APP_SECRET"))

	application.GRPCServer.MustRun()
	log.Info("db connected")
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
