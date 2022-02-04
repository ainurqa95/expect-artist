package seeders

import (
	"github.com/ainurqa95/expect-artist/pkg/entities"
)

func (seed *Seed) SeedArtists() error {
	existsData := seed.repositories.ArtistRepository.ExistsData()

	if existsData {
		return nil
	}

	return seed.saveArtists()
}

func (seed *Seed) saveArtists() error {
	artists := seed.getArtists()
	for _, artist := range artists {
		entity := entities.Artist{
			Name:        artist["name"].(string),
			Description: artist["description"].(string),
			Sorting:     0,
		}
		genres := artist["genres"].([]string)

		genresEntites, err := seed.repositories.GenreRepository.GetByCodes(genres)
		if err != nil {
			return err
		}
		_, err = seed.repositories.ArtistRepository.Create(entity, genresEntites)
		if err != nil {
			return err
		}
	}

	return nil
}

func (seed *Seed) getArtists() []map[string]interface{} {
	// TODO load from source
	return []map[string]interface{}{
		{
			"name": "Король и шут", "description": "Король и Шут» (сокращённо «КиШ») — российская хоррор-панк-группа из Санкт-Петербурга.",
			"genres": []string{"punk", "russian_rock"},
		},
		{
			"name": "Би - 2", "description": "Би-2 — советская, далее российская рок-группа, образованная в 1988 году в Бобруйске.",
			"genres": []string{"russian_rock"},
		},
		{
			"name": "Порнофильмы", "description": "Порнофильмы» — российская панк-рок группа из города Дубна, основанная в 2008 году.",
			"genres": []string{"punk", "russian_rock"},
		},
		{
			"name": "Элизиум", "description": "Эли́зиум» — российская рок-группа из Нижнего Новгорода, основанная её бас-гитаристом и лидером Дмитрием Кузнецовым в 1995 году. ",
			"genres": []string{"punk", "russian_rock"},
		},
		{
			"name": "Люмен (Lumen)", "description": "Российская рок-группа из Уфы, основанная в 1998 году ",
			"genres": []string{"alternative", "russian_rock"},
		},
		{
			"name": "Каста", "description": "Ка́ста» — российская рэп-группа из Ростова-на-Дону, артисты лейбла Respect Production. В состав группы входят Влади, Шым, Хамиль, Змей. ",
			"genres": []string{"rap"},
		},
	}
}
