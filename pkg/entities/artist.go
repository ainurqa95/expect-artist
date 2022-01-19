package entities

const ArtistTable = "artists"

type Artist struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
	Sorting     int    `json:"sorting" db:"sorting"`
}
