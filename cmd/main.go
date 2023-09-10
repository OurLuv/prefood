package main

import (
	"fmt"
	"log"

	"github.com/OurLuv/prefood/internal/config"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}