package seeders

import (
	"github.com/ainurqa95/expect-artist/pkg/repositories"
	"github.com/ainurqa95/expect-artist/pkg/services"
)

type Seed struct {
	repositories *repositories.Repository
	services     *services.Service
}

func NewSeed(repositories *repositories.Repository, services *services.Service) *Seed {
	return &Seed{
		repositories: repositories,
		services:     services,
	}
}

func (seed *Seed) SeedData() error {
	err := seed.SeedCities()
	err = seed.SeedGenres()
	err = seed.SeedArtists()
	err = seed.SeedMessageTypes()
	err = seed.SeedEvents()

	return err
}
