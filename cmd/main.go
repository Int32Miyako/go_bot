package main

import (
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

	err = tgBot.Polling()
	if err != nil {
		panic(err)
	}

}
