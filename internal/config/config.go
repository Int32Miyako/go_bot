package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken             string
	OpenWeatherMapAPIKey string
}

// LoadConfig загружает конфигурацию из .env файла и
// устанавливает переменные окружения.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, errors.New("BOT_TOKEN environment variable not set")
	}

	openWeatherApiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if openWeatherApiKey == "" {
		return nil, errors.New("OPEN_WEATHER_API_KEY environment variable not set")
	}
	return &Config{
		BotToken:             token,
		OpenWeatherMapAPIKey: openWeatherApiKey,
	}, nil
}
