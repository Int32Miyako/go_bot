package telegram

import (
	"log/slog"

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

func (a *Adapter) Polling() error {

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
		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := a.bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}

	}

	return nil
}
