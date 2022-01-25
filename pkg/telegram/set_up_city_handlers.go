package telegram

import (
	"strconv"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) handleSearchCity(message *tgbotapi.Message) error {

	artistName := message.Text

	artistsForSelect, err := bot.services.ArtistManager.SearchArtistByName(artistName)

	if len(artistsForSelect) == 0 || err != nil {
		return artistNotFound
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, bot.messages.Responses.SelectArtist)

	msg.ReplyMarkup = bot.defineOptionsForSelect(artistsForSelect)

	if _, err = bot.client.Send(msg); err != nil {
		return err
	}

	_, err = bot.saveMessageToStorage(message.Chat.ID, artistName, entities.AfterSearchArtistCommand)

	return err
}

func (bot *Bot) handleSetUpCity(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	user, err := bot.saveMessageToStorage(chatId, callbackQuery.Data, entities.ChosedArtist)
	if err != nil {
		return err
	}

	artistId, err := strconv.Atoi(callbackQuery.Data)
	if err != nil {
		return err
	}

	message, err := bot.subscribe(artistId, user.Id)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, message)

	if _, err = bot.client.Send(msg); err != nil {
		return err
	}

	return nil
}
