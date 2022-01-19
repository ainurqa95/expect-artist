package entities

const GenresTable = "genres"

type GenresCodes []string

type Genre struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
	Code string `json:"code" db:"code" binding:"required"`
}
