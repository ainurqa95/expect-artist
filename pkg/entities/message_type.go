package entities

const (
	AfterSearchArtistType = "search_artist"
	AfterSetUpCityType    = "set_up_city"
	CommadType            = "command"
)

const TelegramMessageTypeTable = "telegram_message_types"

type TelegramMessageType struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
	Code string `json:"code" db:"code" binding:"required"`
}
