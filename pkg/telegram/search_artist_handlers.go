package telegram

import (
	"strconv"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) handleSearchArtist(message *tgbotapi.Message) error {

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

	_, err = bot.services.MessageManager.SaveUserMessageToStorage(message.Chat.ID, artistName, entities.AfterSearchArtistCommand)

	return err
}

func (bot *Bot) defineOptionsForSelect(artists []entities.Artist) tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	for _, artist := range artists {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(artist.Name, strconv.Itoa(artist.Id)))
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttons...),
	)
}

func (bot *Bot) handleSubsribeArtist(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	user, err := bot.services.MessageManager.SaveUserMessageToStorage(chatId, callbackQuery.Data, entities.ChosedArtist)
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

func (bot *Bot) subscribe(artistId int, userId int) (string, error) {

	exists := bot.services.SubscriptionManager.ExistsSubscriptionBy(artistId, userId)

	if exists {
		return bot.messages.Responses.SubscriptionAlreadyExists, nil
	}

	err := bot.services.SubscriptionManager.Create(artistId, userId)

	if err != nil {
		return "", err
	}

	return bot.messages.Responses.SubscriptionSuccesfullyAdded, nil
}
