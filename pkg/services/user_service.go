package services

import (
	"database/sql"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) FindOrCreateUser(telegramId int64) (entities.User, error) {
	findedUser, err := userService.userRepository.FindByTelegramId(telegramId)

	if findedUser.Id != 0 && err == nil {
		return findedUser, nil
	}
	convertedTelegramId := sql.NullInt64{Int64: telegramId, Valid: true}
	toCreateUser := entities.User{TelegramId: convertedTelegramId}

	id, err := userService.userRepository.Create(toCreateUser)
	if id != 0 {
		toCreateUser.SetId(id)
	}

	return toCreateUser, err
}

func (userService *UserService) FindUserSubscriptions(user entities.User) ([]entities.Artist, error) {
	return userService.userRepository.FindUserSubscriptions(user)
}

func (userService *UserService) UpdateCity(user entities.User, cityId int) error {
	user.CityId = sql.NullInt32{Int32: int32(cityId), Valid: true}

	return userService.userRepository.Update(user)
}
