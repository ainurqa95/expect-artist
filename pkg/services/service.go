package services

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type UserManager interface {
	FindOrCreateUser(telegramId int64) (entities.User, error)
	FindUserSubscriptions(user entities.User) ([]entities.Artist, error)
	UpdateCity(user entities.User, cityId int) error
}

type MessageManager interface {
	SaveUserMessageToStorage(chatId int64, textMessage string, messageTypeCode string) (entities.User, error)
	FindLastMessage(chatId int64) (entities.Message, error)
}

type ArtistManager interface {
	SearchArtistByName(artistName string) ([]entities.Artist, error)
}

type SubscriptionManager interface {
	ExistsSubscriptionBy(artistId int, userId int) bool
	Create(artistId int, userId int) error
}

type CityManager interface {
	SearchCityByName(cityName string) ([]entities.City, error)
}

type EventManager interface {
	FindComingEvents(cityId int, artists []entities.Artist) ([]entities.Event, error)
	PurifyEventsToShow(events []entities.Event) []string
	Create(event entities.Event) (int, error)
}

type Service struct {
	UserManager
	MessageManager
	ArtistManager
	SubscriptionManager
	CityManager
	EventManager
}

func NewService(repos *repositories.Repository) *Service {
	userService := NewUserService(repos.UserRepository)
	return &Service{
		UserManager:         userService,
		MessageManager:      NewMessageService(userService, repos.MessageRepository, repos.MessageTypeRepository),
		ArtistManager:       NewArtistService(repos.ArtistRepository),
		SubscriptionManager: NewSubscriptionService(repos.SubscriptionRepository),
		CityManager:         NewCityService(repos.CityRepository),
		EventManager:        NewEventService(repos.EventRepository),
	}
}
