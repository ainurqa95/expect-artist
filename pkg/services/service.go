package services

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type UserManager interface {
	FindOrCreateUser(telegramId int64) (entities.User, error)
}

type MessageManager interface {
	SaveMessage(messageTypeCode string, chatId int64, userId int, text string) (entities.Message, error)
	FindLastMessage(chatId int64) (entities.Message, error)
}

type ArtistManager interface {
	SearchArtistByName(artistName string) ([]entities.Artist, error)
}

type SubscriptionManager interface {
	ExistsSubscriptionBy(artistId int, userId int) bool
	Create(artistId int, userId int) error
}

type Service struct {
	UserManager
	MessageManager
	ArtistManager
	SubscriptionManager
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		UserManager:         NewUserService(repos.UserRepository),
		MessageManager:      NewMessageService(repos.MessageRepository, repos.MessageTypeRepository),
		ArtistManager:       NewArtistService(repos.ArtistRepository),
		SubscriptionManager: NewSubscriptionService(repos.SubscriptionRepository),
	}
}
