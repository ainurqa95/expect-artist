package entities

const (
	SearchArtistCommand      = "search_artist_comand"
	AfterSearchArtistCommand = "after_search_artist_comand"
	ChosedArtist             = "chosed_artist"
	SetUpCityCommand         = "set_up_city_command"
	EventsCommand            = "events_command"
	AfterSetUpCityCommand    = "after_set_up_city_command"
	ChosedCity               = "chosed_city"
	OtherMessageCommand      = "other_command_or_message"
)

const MessageTypeTable = "telegram_message_types"

type MessageType struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
	Code string `json:"code" db:"code" binding:"required"`
}
