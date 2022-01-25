package entities

const SubscriptionTable = "subscriptions"

type Subscription struct {
	ArtistId int `json:"artist_id" db:"artist_id"`
	UserId   int `json:"user_id" db:"user_id"`
}
