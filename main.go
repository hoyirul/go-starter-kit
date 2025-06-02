package main

import (
	"fmt"

	"github.com/hoyirul/go-starter-kit/config"
	"github.com/hoyirul/go-starter-kit/pkg/logger"
	"github.com/hoyirul/go-starter-kit/routes"

	"log"
	"os"
)

func main() {
	logger.Init()

	config.LoadEnv()
	config.InitTimezone()
	config.InitDB()
	config.InitRedis()

	defer config.CloseDB()
	defer config.CloseRedis()

	r := routes.SetupRouter()

	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf("%s:%s", host, port)

	log.Printf("Starting server at %s\n", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

/*
List untuk besok yang harus dikerjakan:
	1. Auth JWT √
	2. Middleware √
	3. Swagger
	4. Logging √
	5. Error Handling
	6. Unit Testing √
	7. Integration Testing
	8. Rate Limiting
	9. CORS
	10. Generator untuk CRUD
*/