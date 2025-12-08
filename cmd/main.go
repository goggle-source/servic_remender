package main

import (
	"log/slog"
	"os"
	"servic_remender/internal/config"
)

const (
	envLocal = "lcoal"
	envProd  = "prod"
)

func main() {
	cfg := config.Load()

	log := SetupLogger(cfg.Env)

	log.Info("starting server", slog.Any("grpc", cfg.GRPC))
}

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
