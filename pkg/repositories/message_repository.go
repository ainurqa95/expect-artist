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

func NewDbTelegramMessageRepository(baseDbRepositroy *BaseDbRepository) *DbMessageTypeRepository {
	return &DbMessageTypeRepository{
		db:             baseDbRepositroy.db,
		baseRepository: baseDbRepositroy,
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
