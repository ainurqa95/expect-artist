package services

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type MessageService struct {
	messageRepository     repositories.MessageRepository
	messsgeTypeRepository repositories.MessageTypeRepository
}

func NewMessageService(messageRepository repositories.MessageRepository, messageTypeRepository repositories.MessageTypeRepository) *MessageService {
	return &MessageService{messageRepository: messageRepository, messsgeTypeRepository: messageTypeRepository}
}

func (messageService *MessageService) SaveMessage(messageTypeCode string, chatId int64, userId int) (entities.Message, error) {
	// messageType, err := messageService.messsgeTypeRepository.FindByCode(messageTypeCode)
	// message := entites.Message{}

}
