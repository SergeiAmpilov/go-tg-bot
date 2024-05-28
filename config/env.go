package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
	StickerId string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func New() *Config {
	return &Config{
		Token: getEnv("TOKEN", ""),
		StickerId: getEnv("STICKER_ID", "CAACAgIAAxkBAAEMNjFmVg4GCllRKjUHAwb85hqmMbdnRgACZwMAAhM5jxFB5uPvfPHZUDUE"),
	}
}

func getEnv(key string, defaultValue string) string {
	val, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	return val
}
