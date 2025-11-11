package telegram

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
