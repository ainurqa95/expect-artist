package services

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type ArtistService struct {
	artistRepository repositories.ArtistRepository
}

func NewArtistService(artistRepository repositories.ArtistRepository) *ArtistService {
	return &ArtistService{artistRepository: artistRepository}
}

func (artistService *ArtistService) SearchArtistByName(artistName string) ([]entities.Artist, error) {
	return artistService.artistRepository.SearchArtistByName(artistName)
}
