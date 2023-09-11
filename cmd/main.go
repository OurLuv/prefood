package main

import (
	"fmt"
	"log"
	"os"

	"github.com/OurLuv/prefood/internal/config"
	"github.com/OurLuv/prefood/internal/storage/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

// init is invoked before main()
func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	//* config
	initEnv()
	cfg := config.MustLoad()
	fmt.Println(cfg)

	//* logger
	log := setupLogger()
	log.Info("Test info")

	//* storage
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil{
		log.Error("failed to init storage", err)
		os.Exit(1)
	}
	log.Info("Storage init")
	_ = storage
}

func setupLogger() *slog.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}
