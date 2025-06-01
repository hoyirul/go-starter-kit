package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hoyirul/go-starter-kit/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "gorm.io/driver/sqlserver"

	"github.com/hoyirul/go-starter-kit/config"
)

func main() {
	logger.Init()
	config.LoadEnv()

	action := "up"
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	dbType := os.Getenv("DB_CONN")
	var dsn string
	var driver string

	switch dbType {
	case "sqlite":
		driver = "sqlite3"
		dsn = os.Getenv("DB_PATH")
	case "mysql":
		driver = "mysql"
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	case "pgsql", "postgres":
		driver = "postgres"
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	case "mssql":
		driver = "sqlserver"
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	default:
		logger.LogError("Unsupported database type", fmt.Errorf("unsupported DB_CONN: %s", dbType))
		return
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		logger.LogError("Failed to connect to database", err)
		return
	}
	defer db.Close()

	sqlFile := fmt.Sprintf("migrations/%s/%s.sql", dbType, action)
	content, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		logger.LogError(fmt.Sprintf("Failed to read migration file %s", sqlFile), err)
		return
	}

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		_, err := db.Exec(query)
		if err != nil {
			logger.LogError(fmt.Sprintf("Failed to execute query: %s", query), err)
			return
		} else {
			logger.LogInfo(fmt.Sprintf("Executed query: %s", query))
		}
	}

	logger.LogInfo(fmt.Sprintf("Migration %s completed successfully", action))
}
