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
	Create(messageType entities.MessageType) (int, error)
	FindByCode(code string) (entities.MessageType, error)
	ExistsData() bool
}

type MessageRepository interface {
	Create(message entities.Message) (int, error)
}

type UserRepository interface {
	Create(user entities.User) (int, error)
	FindByTelegramId(telegramId int64) (entities.User, error)
}

type Repository struct {
	CityRepository
	GenreRepository
	ArtistRepository
	MessageTypeRepository
	MessageRepository
	UserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	baseDbRepository := NewBaseDbRepository(db)
	return &Repository{
		CityRepository:        NewDbCityRepository(baseDbRepository),
		GenreRepository:       NewDbGenreRepository(baseDbRepository),
		ArtistRepository:      NewDbArtistRepository(baseDbRepository),
		MessageTypeRepository: NewDbMessageTypeRepository(baseDbRepository),
		UserRepository:        NewDbUserRepository(baseDbRepository),
	}
}
