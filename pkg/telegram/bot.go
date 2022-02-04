package telegram

import (
	"github.com/ainurqa95/expect-artist/pkg/config"
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCommand     = "start"
	searchCommand    = "search"
	setUpCityCommand = "set_up_city"
	eventsCommand    = "events"
	helpCommand      = "help"
)

type Bot struct {
	client   *tgbotapi.BotAPI
	services *services.Service
	messages config.Messages
}

func NewBot(
	client *tgbotapi.BotAPI,
	services *services.Service,
	messages config.Messages,
) *Bot {
	return &Bot{
		client:   client,
		messages: messages,
		services: services,
	}
}

func (bot *Bot) Start() {
	updates := bot.initUpdates()

	bot.handleUpdates(updates)
}

func (bot *Bot) initUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.client.GetUpdatesChan(u)
}

func (bot *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			err := bot.handleCallback(update.CallbackQuery)
			if err != nil {
				bot.handleError(update.CallbackQuery.From.ID, err)
			}
			continue
		}
		if update.Message.IsCommand() {
			err := bot.handleCommand(update.Message)
			if err != nil {
				bot.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		err := bot.handleMessage(update.Message)
		if err != nil {
			bot.handleError(update.Message.Chat.ID, err)
		}
	}
}

func (bot *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return bot.handleStartCommand(message)
	case searchCommand:
		return bot.handleSearchCommand(message)
	case setUpCityCommand:
		return bot.handleSetUpCityCommand(message)
	case eventsCommand:
		return bot.handleEventsCommand(message)
	case helpCommand:
		return bot.handleStartCommand(message)

	default:
		return bot.handleUnknownCommandOrMessage(message)
	}
}

func (bot *Bot) handleMessage(message *tgbotapi.Message) error {
	chatId := message.Chat.ID
	previousMessage, err := bot.services.MessageManager.FindLastMessage(chatId)
	if err != nil {
		return err
	}
	switch previousMessage.MessageType.Code {
	case entities.SearchArtistCommand:
		return bot.handleSearchArtist(message)
	case entities.SetUpCityCommand:
		return bot.handleSearchCity(message)
	default:
		return bot.handleUnknownCommandOrMessage(message)
	}
}

func (bot *Bot) handleCallback(callbackQuery *tgbotapi.CallbackQuery) error {
	previousMessage, err := bot.services.MessageManager.FindLastMessage(callbackQuery.From.ID)
	if err != nil {
		return err
	}
	switch previousMessage.MessageType.Code {
	case entities.AfterSearchArtistCommand:
		return bot.handleSubsribeArtist(callbackQuery)
	case entities.AfterSetUpCityCommand:
		return bot.handleSetUpCity(callbackQuery)
	case entities.ChosedArtist:
		return bot.handleSubsribeArtist(callbackQuery)
	default:
		return callbackHandlerNotFound
	}
}
