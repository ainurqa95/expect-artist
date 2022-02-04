package entities

import "database/sql"

const EventsTable = "events"

const HappenAtDateFormat = "2006-01-02 15:04:05"

type Event struct {
	Id          int            `json:"id" db:"id"`
	Title       string         `json:"title" db:"title"`
	PlaceName   sql.NullString `json:"place_name" db:"place_name"`
	Description string         `json:"description" db:"description"`
	HappenAt    string         `json:"happen_at" db:"happen_at"`
	BuyLink     sql.NullString `json:"buy_link" db:"buy_link"`
	CityId      int            `json:"city_id" db:"city_id"`
	ArtistId    int            `json:"artist_id" db:"artist_id"`
	Artist      Artist         `db:"art"`
}
