package telegram

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCommand     = "start"
	searchCommand    = "search"
	setUpCityCommand = "set_up_city"
)

func (bot *Bot) handleStartCommand(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	textMessage := bot.messages.Start
	msg := tgbotapi.NewMessage(chatId, textMessage)

	if _, err := bot.client.Send(msg); err != nil {
		return err
	}
	_, err := bot.saveMessageToStorage(chatId, textMessage, entities.OtherMessageCommand)

	return err
}

func (bot *Bot) handleSearchCommand(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	textMessage := bot.messages.Responses.Search
	msg := tgbotapi.NewMessage(chatId, textMessage)

	if _, err := bot.client.Send(msg); err != nil {
		return err
	}

	_, err := bot.saveMessageToStorage(chatId, textMessage, entities.SearchArtistCommand)

	return err
}

func (bot *Bot) handleSetUpCityCommand(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	textMessage := bot.messages.Responses.SetUpCity
	msg := tgbotapi.NewMessage(chatId, textMessage)

	if _, err := bot.client.Send(msg); err != nil {
		return err
	}

	_, err := bot.saveMessageToStorage(chatId, textMessage, entities.SetUpCityCommand)

	return err
}

func (bot *Bot) handleUnknownCommandOrMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, bot.messages.UnknownCommand)
	if _, err := bot.client.Send(msg); err != nil {
		return err
	}

	_, err := bot.saveMessageToStorage(message.Chat.ID, bot.messages.UnknownCommand, entities.OtherMessageCommand)

	return err
}
