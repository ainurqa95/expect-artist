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

func (repository *DbUserRepository) Create(user entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, telegram_id, city_id) values ($1, $2, $3, $4) RETURNING id", entities.UsersTable)
	fmt.Println("telegram id", user.TelegramId)
	row := repository.db.QueryRow(query, user.Email, user.PasswordHash, user.TelegramId, user.CityId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
