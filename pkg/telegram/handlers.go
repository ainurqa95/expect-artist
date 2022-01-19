package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCommand = "start"
)

func (bot *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return bot.handleStartCommand(message)
	default:
		return bot.handleUnknownCommand(message)
	}
}

func (bot *Bot) handleStartCommand(message *tgbotapi.Message) error {
	// numericKeyboard := tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
	// 		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
	// 		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	// 	),
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
	// 		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
	// 		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	// 	),
	// )

	msg := tgbotapi.NewMessage(message.Chat.ID, bot.messages.Start)

	// switch message.Text {
	// case "open":
	// msg.ReplyMarkup = numericKeyboard

	// }
	_, err := bot.client.Send(msg)

	return err
}

func (bot *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, bot.messages.UnknownCommand)
	_, err := bot.client.Send(msg)

	return err
}

func (bot *Bot) handleMessage(message *tgbotapi.Message) error {

	return nil
}
