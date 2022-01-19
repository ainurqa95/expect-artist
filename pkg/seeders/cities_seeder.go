package seeders

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
)

func (seed *Seed) SeedCities() error {
	existsData := seed.repositories.CityRepository.ExistsData()

	if existsData {
		return nil
	}

	return seed.saveCities()
}

func (seed *Seed) saveCities() error {
	cities := seed.getCities()
	for _, cityName := range cities {
		cityEntity := entities.City{
			Name: cityName,
		}
		_, err := seed.repositories.CityRepository.Create(cityEntity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (seed *Seed) getCities() []string {
	// TODO load from source
	return []string{"Москва", "Казань", "Санкт-Петербург", "Новосибирск", "Екатеринбург", "Нижний Новгород", "Челябинск", "Самара", "Омск"}
}
