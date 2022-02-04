package entities

import "database/sql"

const UsersTable = "users"

type User struct {
	Id           int            `json:"id" db:"id"`
	Email        sql.NullString `json:"email" db:"email"`
	PasswordHash sql.NullString `json:"password_hash" db:"password_hash"`
	TelegramId   sql.NullInt64  `json:"telegram_id" db:"telegram_id"`
	CreatedAt    sql.NullString `json:"created_at" db:"created_at"`
	CityId       sql.NullInt32  `json:"city_id" db:"city_id"`
}

func (user *User) SetId(id int) {
	user.Id = id
}

func (user *User) GetCityId() int {
	return int(user.CityId.Int32)
}
