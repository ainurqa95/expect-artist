package repositories

import (
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type DbMessageRepository struct {
	db             *sqlx.DB
	baseRepository *BaseDbRepository
}

func NewDbMessageRepository(baseDbRepository *BaseDbRepository) *DbMessageRepository {
	return &DbMessageRepository{
		db:             baseDbRepository.db,
		baseRepository: baseDbRepository,
	}
}

func (repository *DbMessageRepository) Create(message entities.Message) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (chat_id, message, message_type_id, user_id) values ($1, $2, $3, $4) RETURNING id", entities.MessageTable)
	row := repository.db.QueryRow(query, message.ChatId, message.Message, message.MessageTypeId, message.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *DbMessageRepository) FindLastMessage(chatId int64) (entities.Message, error) {
	var message entities.Message
	query := fmt.Sprintf(`SELECT tg.*, tmt.code "tmt.code", tmt.name "tmt.name"
	 FROM %s tg JOIN %s tmt ON tg.message_type_id = tmt.id  WHERE tg.chat_id=$1 ORDER BY tg.id DESC LIMIT 1;`, entities.MessageTable, entities.MessageTypeTable)
	err := repository.db.Get(&message, query, chatId)

	return message, err
}
