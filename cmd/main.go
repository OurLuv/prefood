package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OurLuv/prefood/internal/config"
	"github.com/OurLuv/prefood/internal/server/handler"
	"github.com/OurLuv/prefood/internal/service"
	"github.com/OurLuv/prefood/internal/storage/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

// @title PreFood
// @version 1.0
// @description REST API server for managing data in a food ordering application.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

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
	dbPath, exists := os.LookupEnv("DB_PATH")
	if !exists {
		log.Error("db path is not set: %s", dbPath)
	}
	pool, err := postgres.NewDB(dbPath)
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}
	defer pool.Close()
	log.Info("Storage init")

	//* initing repos, services & handlers
	repos := postgres.NewRepository(pool)
	log.Info("repos inited")
	services := service.NewService(*repos)
	log.Info("services inited")
	handlers := handler.NewHandler(*services, log)
	log.Info("handlers inited")

	//* starting server
	router := handlers.InitRoutes()
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}
	log.Info("server is started")
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start a server", err)
	}
}

func setupLogger() *slog.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}
