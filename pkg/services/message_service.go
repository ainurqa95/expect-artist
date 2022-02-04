package services

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type MessageService struct {
	userManager           UserManager
	messageRepository     repositories.MessageRepository
	messageTypeRepository repositories.MessageTypeRepository
}

func NewMessageService(
	userManager UserManager,
	messageRepository repositories.MessageRepository,
	messageTypeRepository repositories.MessageTypeRepository,
) *MessageService {
	return &MessageService{
		userManager:           userManager,
		messageRepository:     messageRepository,
		messageTypeRepository: messageTypeRepository,
	}
}

func (messageService *MessageService) SaveUserMessageToStorage(chatId int64, textMessage string, messageTypeCode string) (entities.User, error) {
	user, err := messageService.userManager.FindOrCreateUser(chatId)
	if err != nil {
		return user, err
	}

	_, err = messageService.saveMessage(messageTypeCode, chatId, user.Id, textMessage)

	return user, err
}

func (messageService *MessageService) saveMessage(messageTypeCode string, chatId int64, userId int, text string) (entities.Message, error) {
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
