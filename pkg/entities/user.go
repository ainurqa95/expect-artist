package entities

const UsersTable = "users"

type User struct {
	Id          int    `json:"id" db:"id"`
	Email       string `json:"email" db:"name"`
	PasswordHas string `json:"password_hash" db:"password_hash"`
	TelegramId  int64  `json:"telegram_id" db:"telegram_id"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	CityId      int    `json:"city_id" db:"city_id"`
}
