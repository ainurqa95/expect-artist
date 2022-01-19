package seeders

import (
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

type Seed struct {
	repositories *repositories.Repository
}

func NewSeed(repositories *repositories.Repository) *Seed {
	return &Seed{
		repositories: repositories,
	}
}

func (seed *Seed) SeedData() error {
	err := seed.SeedCities()
	err = seed.SeedGenres()
	err = seed.SeedArtists()
	err = seed.SeedMessageTypes()

	return err
}
