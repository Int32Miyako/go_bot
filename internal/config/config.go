package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig загружает конфигурацию из .env файла и
// устанавливает переменные окружения.
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("BOT_TOKEN environment variable not set")
	}

}
