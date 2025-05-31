package main

import (
	"fmt"
	"go-starter-kit/pkg/logger"
	"os"
	"strings"

	"go-starter-kit/config"
	"go-starter-kit/seeders"
)

func main() {
	logger.Init()
	
	if len(os.Args) < 2 {
		logger.LogError("No command provided. Use 'seed' or 'unseed'.", nil)
		os.Exit(1)
	}

	config.LoadEnv()
	config.InitDB()
	defer config.CloseDB()

	db := config.DB

	userSeeder := &seeders.UserSeeder{}
	productSeeder := &seeders.ProductSeeder{}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "seed":
		if err := userSeeder.Seed(db); err != nil {
			logger.LogError("Failed to seed users", err)
		}
		if err := productSeeder.Seed(db); err != nil {
			logger.LogError("Failed to seed products", err)
			os.Exit(1)
		}
		logger.LogInfo("Seeding completed successfully.")

	case "unseed":
		if err := productSeeder.Unseed(db); err != nil {
			logger.LogError("Failed to unseed products", err)
			os.Exit(1)
		}
		if err := userSeeder.Unseed(db); err != nil {
			logger.LogError("Failed to unseed users", err)
			os.Exit(1)
		}
		logger.LogInfo("Unseeding completed successfully.")

	default:
		logger.LogError(fmt.Sprintf("Unknown command: %s", cmd), nil)
		return
	}
}
