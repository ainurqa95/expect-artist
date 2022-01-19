package entities

const CitiesTable = "cities"

type City struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}
