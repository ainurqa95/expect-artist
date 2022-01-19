package repositories

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type CityRepository interface {
	Create(city entities.City) (int, error)
	ExistsData() bool
}

type GenreRepository interface {
	Create(genre entities.Genre) (int, error)
	GetByCodes(codes []string) ([]entities.Genre, error)
	ExistsData() bool
}

type ArtistRepository interface {
	Create(artist entities.Artist, genres []entities.Genre) (int, error)
	ExistsData() bool
}

type MessageTypeRepository interface {
	Create(messageType entities.TelegramMessageType) (int, error)
	ExistsData() bool
}

type Repository struct {
	CityRepository
	GenreRepository
	ArtistRepository
	MessageTypeRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	baseDbRepository := NewBaseDbRepository(db)
	return &Repository{
		CityRepository:        NewDbCityRepository(baseDbRepository),
		GenreRepository:       NewDbGenreRepository(baseDbRepository),
		ArtistRepository:      NewDbArtistRepository(baseDbRepository),
		MessageTypeRepository: NewDbMessageTypeRepository(baseDbRepository),
	}
}
