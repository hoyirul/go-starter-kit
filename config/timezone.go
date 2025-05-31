package config

import (
	"log"
	"os"
	"time"
)

func InitTimezone() {
	tz := os.Getenv("TIMEZONE")
	if tz == "" {
		tz = "Asia/Jakarta"
	}

	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatalf("Failed to load timezone '%s': %v", tz, err)
	}

	time.Local = location
	log.Printf("Timezone set to: %s", location)
}
