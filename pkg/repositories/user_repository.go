package repositories

import (
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type DbUserRepository struct {
	db             *sqlx.DB
	baseRepository *BaseDbRepository
}

func NewDbUserRepository(baseRepository *BaseDbRepository) *DbUserRepository {
	return &DbUserRepository{db: baseRepository.db, baseRepository: baseRepository}
}

func (repository *DbUserRepository) FindByTelegramId(telegramId int64) (entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE telegram_id = $1 limit 1", entities.UsersTable)
	err := repository.db.Get(&user, query, telegramId)

	return user, err
}

func (repository *DbUserRepository) FindUserSubscriptions(user entities.User) ([]entities.Artist, error) {
	var artists []entities.Artist

	query := fmt.Sprintf("SELECT  art.id, art.name, art.description, art.sorting FROM %s sbscr join %s art ON art.id = sbscr.artist_id WHERE sbscr.user_id = $1", entities.SubscriptionTable, entities.ArtistTable)
	err := repository.db.Select(&artists, query, user.Id)

	return artists, err
}

func (repository *DbUserRepository) Create(user entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, telegram_id, city_id) VALUES ($1, $2, $3, $4) RETURNING id", entities.UsersTable)
	row := repository.db.QueryRow(query, user.Email, user.PasswordHash, user.TelegramId, user.CityId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *DbUserRepository) Update(user entities.User) error {

	query := fmt.Sprintf("UPDATE %s SET email=$1, password_hash=$2, telegram_id=$3, city_id=$4 where id=$5", entities.UsersTable)
	_, err := repository.db.Exec(query, user.Email, user.PasswordHash, user.TelegramId, user.CityId, user.Id)

	return err

}
