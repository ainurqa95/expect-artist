package repositories

import (
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type DbMessageTypeRepository struct {
	db             *sqlx.DB
	baseRepository *BaseDbRepository
}

func NewDbMessageTypeRepository(baseRepository *BaseDbRepository) *DbMessageTypeRepository {
	return &DbMessageTypeRepository{db: baseRepository.db, baseRepository: baseRepository}
}

func (repository *DbMessageTypeRepository) Create(messageType entities.TelegramMessageType) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, code) values ($1, $2) RETURNING id", entities.TelegramMessageTypeTable)

	row := repository.db.QueryRow(query, messageType.Name, messageType.Code)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *DbMessageTypeRepository) ExistsData() bool {
	return repository.baseRepository.ExistsData(entities.TelegramMessageTypeTable)
}
