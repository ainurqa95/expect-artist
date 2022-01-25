package entities

const ArtistGenreTable = "artist_genre"

type ArtistGenre struct {
	ArtistId int `json:"artist_id" db:"artist_id"`
	GenreId  int `json:"genre_id" db:"genre_id"`
}
