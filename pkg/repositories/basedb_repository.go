package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BaseDbRepository struct {
	db *sqlx.DB
}

func NewBaseDbRepository(db *sqlx.DB) *BaseDbRepository {
	return &BaseDbRepository{db: db}
}

func (baseRepository *BaseDbRepository) ExistsData(table string) bool {
	var exists int
	query := fmt.Sprintf("SELECT 1 FROM %s", table)

	err := baseRepository.db.Get(&exists, query)

	return exists == 1 && err == nil
}
