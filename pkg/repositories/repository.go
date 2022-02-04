package repositories

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type CityRepository interface {
	Create(city entities.City) (int, error)
	SearchCityByName(cityName string) ([]entities.City, error)
	ExistsData() bool
}

type GenreRepository interface {
	Create(genre entities.Genre) (int, error)
	GetByCodes(codes []string) ([]entities.Genre, error)
	ExistsData() bool
}

type ArtistRepository interface {
	Create(artist entities.Artist, genres []entities.Genre) (int, error)
	SearchArtistByName(artistName string) ([]entities.Artist, error)
	ExistsData() bool
}

type MessageTypeRepository interface {
	Create(messageType entities.MessageType) (int, error)
	FindByCode(code string) (entities.MessageType, error)
	ExistsData() bool
}

type MessageRepository interface {
	Create(message entities.Message) (int, error)
	FindLastMessage(chatId int64) (entities.Message, error)
}

type UserRepository interface {
	Create(user entities.User) (int, error)
	FindByTelegramId(telegramId int64) (entities.User, error)
	FindUserSubscriptions(user entities.User) ([]entities.Artist, error)
	Update(user entities.User) error
}

type SubscriptionRepository interface {
	ExistsSubscriptionBy(artistId int, userId int) bool
	Create(artistId int, userId int) error
}

type EventRepository interface {
	Create(event entities.Event) (int, error)
	FindEventsBy(city entities.City, artists []entities.Artist, maxHappenAt string) ([]entities.Event, error)
	ExistsData() bool
}

type Repository struct {
	CityRepository
	GenreRepository
	ArtistRepository
	MessageTypeRepository
	MessageRepository
	UserRepository
	SubscriptionRepository
	EventRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	baseDbRepository := NewBaseDbRepository(db)
	return &Repository{
		CityRepository:         NewDbCityRepository(baseDbRepository),
		GenreRepository:        NewDbGenreRepository(baseDbRepository),
		ArtistRepository:       NewDbArtistRepository(baseDbRepository),
		MessageTypeRepository:  NewDbMessageTypeRepository(baseDbRepository),
		UserRepository:         NewDbUserRepository(baseDbRepository),
		MessageRepository:      NewDbMessageRepository(baseDbRepository),
		SubscriptionRepository: NewDbSubscriptionRepository(baseDbRepository),
		EventRepository:        NewDbEventRepository(baseDbRepository),
	}
}
