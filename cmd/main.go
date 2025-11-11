package main

import (
	"meteo_bot/internal/adapters/gis_meteo"
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

	gis := gis_meteo.NewGisMeteoAPI()
	err = tgBot.Polling(gis)
	if err != nil {
		panic(err)
	}

}
