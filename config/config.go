package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/hoyirul/go-starter-kit/pkg/logger"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	filename := fmt.Sprintf("envs/.env.%s", env)
	if err := godotenv.Load(filename); err != nil {
		logger.LogError(fmt.Sprintf("Error loading .env file: %s", filename), err)
	}
}

func InitDB() {
	var err error
	dbType := os.Getenv("DB_CONN")

	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	case "pgsql", "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	case "sqlite":
		dsn := os.Getenv("DB_PATH")
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	case "mssql":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	default:
		logger.LogError("Unsupported database type", fmt.Errorf("unsupported db type: %s", dbType))
		return
	}

	if err != nil {
		logger.LogError("Failed to connect to database", err)
		return
	}

	logger.LogInfo("Connected to database successfully")
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		logger.LogError("Failed to get database instance", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		logger.LogError("Failed to close database connection", err)
		return
	}

	logger.LogInfo("Database connection closed successfully")
}