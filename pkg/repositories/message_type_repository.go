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

func (repository *DbMessageTypeRepository) Create(messageType entities.MessageType) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, code) values ($1, $2) RETURNING id", entities.MessageTypeTable)
	row := repository.db.QueryRow(query, messageType.Name, messageType.Code)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *DbMessageTypeRepository) FindByCode(code string) (entities.MessageType, error) {
	var messageType entities.MessageType

	query := fmt.Sprintf("SELECT * FROM %s WHERE code = $1 limit 1", entities.MessageTypeTable)
	err := repository.db.Get(&messageType, query, code)

	return messageType, err
}

func (repository *DbMessageTypeRepository) ExistsData() bool {
	return repository.baseRepository.ExistsData(entities.MessageTypeTable)
}
