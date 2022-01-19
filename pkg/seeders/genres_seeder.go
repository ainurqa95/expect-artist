package seeders

import (
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
)

func (seed *Seed) SeedGenres() error {
	existsData := seed.repositories.GenreRepository.ExistsData()
	fmt.Println(existsData)
	if existsData {
		return nil
	}
	// return nil
	return seed.saveGenres()
}

func (seed *Seed) saveGenres() error {
	genres := seed.getGenres()
	for _, genre := range genres {
		genreEntity := entities.Genre{
			Name: genre["name"],
			Code: genre["code"],
		}
		_, err := seed.repositories.GenreRepository.Create(genreEntity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (seed *Seed) getGenres() []map[string]string {
	// TODO load from source

	return []map[string]string{
		{"code": "punk", "name": "Панк"}, {"code": "rock", "name": "Рок"},
		{"code": "rap", "name": "Рэп и хип-хоп"}, {"code": "electronic", "name": "Электроника"},
		{"code": "russian_rock", "name": "Русский рок"}, {"code": "classic", "name": "Классика"},
		{"code": "jazz", "name": "Джазз"}, {"code": "alternative", "name": "Альтернатива"},
	}

}
