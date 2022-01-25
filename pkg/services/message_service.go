package services

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type MessageService struct {
	messageRepository     repositories.MessageRepository
	messageTypeRepository repositories.MessageTypeRepository
}

func NewMessageService(
	messageRepository repositories.MessageRepository,
	messageTypeRepository repositories.MessageTypeRepository,
) *MessageService {
	return &MessageService{
		messageRepository:     messageRepository,
		messageTypeRepository: messageTypeRepository,
	}
}

func (messageService *MessageService) SaveMessage(messageTypeCode string, chatId int64, userId int, text string) (entities.Message, error) {
	messageType, err := messageService.messageTypeRepository.FindByCode(messageTypeCode)

	message := entities.Message{
		ChatId:        chatId,
		UserId:        userId,
		MessageTypeId: messageType.Id,
		Message:       text,
	}

	id, err := messageService.messageRepository.Create(message)
	if id != 0 {
		message.SetId(id)
	}

	return message, err
}

func (messageService *MessageService) FindLastMessage(chatId int64) (entities.Message, error) {
	return messageService.messageRepository.FindLastMessage(chatId)
}
