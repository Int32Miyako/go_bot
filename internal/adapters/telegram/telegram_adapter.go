package telegram

import (
	"log/slog"
	"meteo_bot/internal/adapters"
	"meteo_bot/internal/adapters/gis_meteo"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Adapter реализует взаимодействие с Telegram Bot API.
type Adapter struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramAdapter(botToken string) (*Adapter, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	return &Adapter{bot: bot}, nil
}

func (a *Adapter) PollingGisMeteo(gis *gis_meteo.GisMeteoAPI) error {
	a.bot.Debug = true
	slog.Info("Authorized on account", "account", a.bot.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := a.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		temperature, err := gis.ServeGisMeteo(strconv.FormatFloat(update.Message.Location.Latitude, 'f', 6, 64),
			strconv.FormatFloat(update.Message.Location.Longitude, 'f', 6, 64))
		if err != nil {
			slog.Error("Error while getting weather data", "error", err)
			msg.Text = "Ошибка при получении данных о погоде."
		} else {
			msg.Text = "Температура: " + temperature + "°C"
		}
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err = a.bot.Send(msg); err != nil {
			slog.Error("Error while sending message", "error", err)
		}

	}

	return nil
}

func (a *Adapter) PollingOpenWeather(meteo *adapters.OpenWeatherAPI) error {
	a.bot.Debug = true
	slog.Info("Authorized on account", "account", a.bot.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := a.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		temperature, err := meteo.ServeOpenWeather(
			strconv.FormatFloat(update.Message.Location.Latitude, 'f', 6, 64),
			strconv.FormatFloat(update.Message.Location.Longitude, 'f', 6, 64))
		if err != nil {
			slog.Error("Error while getting weather data", "error", err)
			msg.Text = "Ошибка при получении данных о погоде."
		} else {
			msg.Text = "Температура: " + temperature + "°C"
		}
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err = a.bot.Send(msg); err != nil {
			slog.Error("Error while sending message", "error", err)
		}

	}

	return nil
}
