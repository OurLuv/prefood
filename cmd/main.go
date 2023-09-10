package main

import (
	"fmt"
	"log"
	"os"

	"github.com/OurLuv/prefood/internal/config"
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
	initEnv()
	cfg := config.MustLoad()
	fmt.Println(cfg)
	log := setupLogger()
	log.Info("Test info")
}

func setupLogger() *slog.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}
