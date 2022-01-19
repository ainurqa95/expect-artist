package entities

const TelegramMessageTable = "telegram_messages"

type TelegramMessage struct {
	Id            int    `json:"id" db:"id"`
	UserId        int    `json:"user_id" db:"user_id" binding:"required"`
	ChatId        int64  `json:"name" db:"name" binding:"required"`
	Message       string `json:"message" db:"message" binding:"required"`
	CreatedAt     string `json:"created_at" db:"created_at"`
	MessageTypeId int    `json:"message_type_id" db:"message_type_id"`
}
