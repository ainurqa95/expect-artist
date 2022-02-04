package telegram

import (
	"strconv"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) handleSearchCity(message *tgbotapi.Message) error {

	cityName := message.Text

	citiesForSelect, err := bot.services.CityManager.SearchCityByName(cityName)

	if len(citiesForSelect) == 0 || err != nil {
		return cityNotFound
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, bot.messages.Responses.SelectArtist)

	msg.ReplyMarkup = bot.defineOptionsForSelectCity(citiesForSelect)

	if _, err = bot.client.Send(msg); err != nil {
		return err
	}

	_, err = bot.services.MessageManager.SaveUserMessageToStorage(message.Chat.ID, cityName, entities.AfterSetUpCityCommand)

	return err
}

func (bot *Bot) defineOptionsForSelectCity(cities []entities.City) tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	for _, city := range cities {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(city.Name, strconv.Itoa(city.Id)))
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttons...),
	)
}

func (bot *Bot) handleSetUpCity(callbackQuery *tgbotapi.CallbackQuery) error {
	chatId := callbackQuery.From.ID
	user, err := bot.services.MessageManager.SaveUserMessageToStorage(chatId, callbackQuery.Data, entities.ChosedCity)
	if err != nil {
		return err
	}

	cityId, err := strconv.Atoi(callbackQuery.Data)
	if err != nil {
		return err
	}

	err = bot.services.UserManager.UpdateCity(user, cityId)

	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, bot.messages.Responses.CitySettingUpIsSuccesfull)

	if _, err = bot.client.Send(msg); err != nil {
		return err
	}

	return nil
}
