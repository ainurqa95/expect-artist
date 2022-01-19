package repositories

import (
	"fmt"
	"strings"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type DbGenreRepository struct {
	db             *sqlx.DB
	baseRepository *BaseDbRepository
}

func NewDbGenreRepository(baseRepository *BaseDbRepository) *DbGenreRepository {
	return &DbGenreRepository{db: baseRepository.db, baseRepository: baseRepository}
}

func (repository *DbGenreRepository) Create(genre entities.Genre) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, code) values ($1, $2) RETURNING id", entities.GenresTable)

	row := repository.db.QueryRow(query, genre.Name, genre.Code)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *DbGenreRepository) GetByCodes(codes []string) ([]entities.Genre, error) {

	implodedGenres := "{"
	implodedGenres = implodedGenres + strings.Join(codes, ",")
	implodedGenres = implodedGenres + "}"
	query := fmt.Sprintf("SELECT id, name, code FROM %s WHERE code = ANY ($1)", entities.GenresTable)
	var genres []entities.Genre
	err := repository.db.Select(&genres, query, implodedGenres)

	return genres, err
}

func (repository *DbGenreRepository) ExistsData() bool {
	return repository.baseRepository.ExistsData(entities.GenresTable)
}
