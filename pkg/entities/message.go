package entities

const MessageTable = "telegram_messages"

type Message struct {
	Id            int        `json:"id" db:"id"`
	UserId        int        `json:"user_id" db:"user_id" binding:"required"`
	ChatId        int64      `json:"chat_id" db:"chat_id" binding:"required"`
	Message       string     `json:"message" db:"message" binding:"required"`
	CreatedAt     string     `json:"created_at" db:"created_at"`
	MessageTypeId int        `json:"message_type_id" db:"message_type_id"`
	MessageType   `db:"tmt"` // see https://snippets.aktagon.com/snippets/757-how-to-join-two-tables-with-jmoiron-sqlx
}

func (message *Message) SetId(id int) {
	message.Id = id
}
