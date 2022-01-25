package repositories

import (
	"database/sql"
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type DbArtistRepository struct {
	db             *sqlx.DB
	baseRepository *BaseDbRepository
}

func NewDbArtistRepository(baseRepository *BaseDbRepository) *DbArtistRepository {
	return &DbArtistRepository{db: baseRepository.db, baseRepository: baseRepository}
}

func (repository *DbArtistRepository) Create(artist entities.Artist, genres []entities.Genre) (int, error) {
	transaction, err := repository.db.Begin()
	if err != nil {
		return 0, err
	}

	var artistId int
	createArtistQuery := fmt.Sprintf("INSERT INTO %s (name, description, sorting) values ($1, $2, $3) RETURNING id", entities.ArtistTable)

	row := transaction.QueryRow(createArtistQuery, artist.Name, artist.Description, artist.Sorting)
	err = row.Scan(&artistId)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	return repository.attachGenres(transaction, artistId, genres)
}

func (repository *DbArtistRepository) attachGenres(transaction *sql.Tx, artistId int, genres []entities.Genre) (int, error) {
	for _, genre := range genres {
		query := fmt.Sprintf("INSERT INTO %s (artist_id, genre_id) values ($1, $2)", entities.ArtistGenreTable)
		_, err := transaction.Exec(query, artistId, genre.Id)
		if err != nil {
			transaction.Rollback()
			return 0, err
		}
	}

	return artistId, transaction.Commit()
}

func (repository *DbArtistRepository) SearchArtistByName(artistName string) ([]entities.Artist, error) {

	query := fmt.Sprintf("SELECT * FROM %s WHERE LOWER(name) like LOWER($1)", entities.ArtistTable)
	var artists []entities.Artist
	err := repository.db.Select(&artists, query, "%"+artistName+"%")

	return artists, err
}

func (repository *DbArtistRepository) ExistsData() bool {
	return repository.baseRepository.ExistsData(entities.ArtistTable)
}
