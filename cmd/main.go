package main

import (
	"meteo_bot/internal/adapters"
	"meteo_bot/internal/adapters/telegram"
	"meteo_bot/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	tgBot, err := telegram.NewTelegramAdapter(cfg.BotToken)
	if err != nil {
		panic(err)
	}

	// gis := gis_meteo.NewGisMeteoAPI()

	openWeather := adapters.NewOpenWeatherAPI(cfg.OpenWeatherMapAPIKey)
	err = tgBot.PollingOpenWeather(openWeather)
	if err != nil {
		panic(err)
	}
}
