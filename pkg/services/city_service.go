package services

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type CityService struct {
	cityRepository repositories.CityRepository
}

func NewCityService(cityRepository repositories.CityRepository) *CityService {
	return &CityService{cityRepository: cityRepository}
}

func (cityService *CityService) SearchCityByName(cityName string) ([]entities.City, error) {
	return cityService.cityRepository.SearchCityByName(cityName)
}
