package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	invalidUrlError = errors.New("Url is invalid")
)

func (bot *Bot) handleError(chatID int64, err error) {
	messageText := err.Error()

	switch err {

	default:
		messageText = bot.messages.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.client.Send(msg)
}
