package repositories

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type DbEventRepository struct {
	db             *sqlx.DB
	baseRepository *BaseDbRepository
}

func NewDbEventRepository(baseRepository *BaseDbRepository) *DbEventRepository {
	return &DbEventRepository{db: baseRepository.db, baseRepository: baseRepository}
}

func (repository *DbEventRepository) Create(event entities.Event) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title, place_name, description, happen_at, buy_link, city_id, artist_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", entities.EventsTable)
	row := repository.db.QueryRow(query, event.Title, event.PlaceName, event.Description, event.HappenAt, event.BuyLink, event.CityId, event.ArtistId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *DbEventRepository) FindEventsBy(city entities.City, artists []entities.Artist, maxHappenAt string) ([]entities.Event, error) {
	var events []entities.Event
	implodedArtists := repository.implodeArtistIds(artists)

	query := fmt.Sprintf(`SELECT evts.*, art.name "art.name", art.description "art.description" FROM %s evts
							JOIN %s art ON evts.artist_id = art.id
							WHERE city_id = $1 AND artist_id = ANY ($2) AND happen_at > now() AND happen_at < $3`, entities.EventsTable, entities.ArtistTable)
	err := repository.db.Select(&events, query, city.Id, implodedArtists, maxHappenAt)

	return events, err
}

func (repository *DbEventRepository) implodeArtistIds(artists []entities.Artist) string {
	implodedArtists := "{"
	var artistIds []string
	for _, artist := range artists {
		artistIds = append(artistIds, strconv.Itoa(artist.Id))
	}
	implodedArtists = implodedArtists + strings.Join(artistIds, ",")
	implodedArtists = implodedArtists + "}"

	return implodedArtists
}

func (repository *DbEventRepository) ExistsData() bool {
	return repository.baseRepository.ExistsData(entities.EventsTable)
}
