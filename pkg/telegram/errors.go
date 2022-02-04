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
	cityHasNotBeenSet       = errors.New("City has not been set")
	artistsHaveNotBeenSet   = errors.New("Artists have not been set")
	eventsNotFound          = errors.New("Events not found")
)

func (bot *Bot) handleError(chatID int64, err error) {
	log.Println(err)
	messageText := err.Error()

	switch err {
	case artistNotFound:
		messageText = bot.messages.Errors.ArtistNotFound
	case cityNotFound:
		messageText = bot.messages.Errors.CityNotFound
	case cityHasNotBeenSet:
		messageText = bot.messages.Errors.CityHasNotBeenSet
	case artistsHaveNotBeenSet:
		messageText = bot.messages.Errors.ArtistsHaveNotBeenSet
	case eventsNotFound:
		messageText = bot.messages.Errors.EventsNotFound
	default:
		messageText = bot.messages.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	bot.client.Send(msg)
}
