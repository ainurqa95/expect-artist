package telegram

import (
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	invalidUrlError         = errors.New("Url is invalid")
	artistNotFound          = errors.New("Artist not found")
	cityNotFound            = errors.New("City not found")
	callbackHandlerNotFound = errors.New("callback not found")
)

func (bot *Bot) handleError(chatID int64, err error) {
	log.Println(err)
	messageText := err.Error()

	switch err {
	case artistNotFound:
		messageText = bot.messages.Errors.ArtistNotFound
	case cityNotFound:
		messageText = bot.messages.Errors.CityNotFound
	default:
		messageText = bot.messages.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.client.Send(msg)
}
