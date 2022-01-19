package services

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type UserManager interface {
	FindOrCreateUser(telegramId int64) (entities.User, error)
}

type MessageManager interface {
	SaveMessage(messageTypeCode string, chatId int64, userId int) (entities.Message, error)
}

type Service struct {
	UserManager
	MessageManager
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		UserManager:    NewUserService(repos.UserRepository),
		MessageManager: NewMessageService(repos.MessageRepository, repos.MessageTypeRepository),
	}
}
