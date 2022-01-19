package repositories

import (
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type DbCityRepository struct {
	db             *sqlx.DB
	baseRepository *BaseDbRepository
}

func NewDbCityRepository(baseRepository *BaseDbRepository) *DbCityRepository {
	return &DbCityRepository{db: baseRepository.db, baseRepository: baseRepository}
}

func (cityRepository *DbCityRepository) Create(city entities.City) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", entities.CitiesTable)

	row := cityRepository.db.QueryRow(query, city.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (cityRepository *DbCityRepository) ExistsData() bool {
	return cityRepository.baseRepository.ExistsData(entities.CitiesTable)
}
