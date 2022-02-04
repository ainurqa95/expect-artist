package telegram

import (
	"log"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) handleStartCommand(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	textMessage := bot.messages.Start
	msg := tgbotapi.NewMessage(chatId, textMessage)

	if _, err := bot.client.Send(msg); err != nil {
		return err
	}
	_, err := bot.services.MessageManager.SaveUserMessageToStorage(chatId, textMessage, entities.OtherMessageCommand)

	return err
}

func (bot *Bot) handleSearchCommand(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	textMessage := bot.messages.Responses.Search
	msg := tgbotapi.NewMessage(chatId, textMessage)

	if _, err := bot.client.Send(msg); err != nil {
		return err
	}

	_, err := bot.services.MessageManager.SaveUserMessageToStorage(chatId, textMessage, entities.SearchArtistCommand)

	return err
}

func (bot *Bot) handleSetUpCityCommand(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	textMessage := bot.messages.Responses.SetUpCity
	msg := tgbotapi.NewMessage(chatId, textMessage)

	if _, err := bot.client.Send(msg); err != nil {
		return err
	}

	_, err := bot.services.MessageManager.SaveUserMessageToStorage(chatId, textMessage, entities.SetUpCityCommand)

	return err
}

func (bot *Bot) handleEventsCommand(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	textMessage := bot.messages.Responses.ShowEvents

	user, err := bot.services.MessageManager.SaveUserMessageToStorage(chatId, textMessage, entities.EventsCommand)
	if err != nil || !user.CityId.Valid {
		return cityHasNotBeenSet
	}
	userArtists, err := bot.services.UserManager.FindUserSubscriptions(user)

	if err != nil || len(userArtists) == 0 {
		return artistsHaveNotBeenSet
	}

	comingEvents, err := bot.services.EventManager.FindComingEvents(user.GetCityId(), userArtists)
	log.Println(comingEvents, err)
	if err != nil || len(comingEvents) == 0 {
		return eventsNotFound
	}

	purifiedEvents := bot.services.EventManager.PurifyEventsToShow(comingEvents)

	return bot.sendPurifiedEvents(chatId, purifiedEvents)
}

func (bot *Bot) sendPurifiedEvents(chatId int64, purifiedEvents []string) error {
	for _, event := range purifiedEvents {
		msg := tgbotapi.NewMessage(chatId, event)
		if _, err := bot.client.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func (bot *Bot) handleUnknownCommandOrMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, bot.messages.UnknownCommand)
	if _, err := bot.client.Send(msg); err != nil {
		return err
	}

	_, err := bot.services.MessageManager.SaveUserMessageToStorage(message.Chat.ID, bot.messages.UnknownCommand, entities.OtherMessageCommand)

	return err
}
