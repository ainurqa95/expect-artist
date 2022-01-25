package repositories

import (
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type DbSubsriptionRepository struct {
	db             *sqlx.DB
	baseRepository *BaseDbRepository
}

func NewDbSubscriptionRepository(baseDbRepository *BaseDbRepository) *DbSubsriptionRepository {
	return &DbSubsriptionRepository{
		db:             baseDbRepository.db,
		baseRepository: baseDbRepository,
	}
}

func (repository *DbSubsriptionRepository) ExistsSubscriptionBy(artistId int, userId int) bool {
	var subscription entities.Subscription

	query := fmt.Sprintf("SELECT * FROM %s WHERE artist_id = $1 and user_id = $2 limit 1", entities.SubscriptionTable)
	err := repository.db.Get(&subscription, query, artistId, userId)

	if err != nil || subscription.ArtistId == 0 || subscription.UserId == 0 {
		return false
	}

	return true
}

func (repository *DbSubsriptionRepository) Create(artistId int, userId int) error {
	subscription := entities.Subscription{
		ArtistId: artistId,
		UserId:   userId,
	}
	query := fmt.Sprintf("INSERT INTO %s (artist_id, user_id) values (:artist_id, :user_id)", entities.SubscriptionTable)
	_, err := repository.db.NamedExec(query, subscription)

	return err
}
