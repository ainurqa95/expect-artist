package seeders

import (
	"database/sql"
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
)

func (seed *Seed) SeedEvents() error {
	existsData := seed.repositories.EventRepository.ExistsData()
	fmt.Println()
	if existsData {
		return nil
	}

	return seed.saveEvents()
}

func (seed *Seed) saveEvents() error {
	events := seed.getEvents()
	for _, event := range events {
		artists, err := seed.services.ArtistManager.SearchArtistByName(event["artist_name"])
		if err != nil {
			return err
		}
		cities, err := seed.services.CityManager.SearchCityByName(event["city_name"])
		if err != nil {
			return err
		}
		artist := artists[0]
		city := cities[0]
		eventEntity := entities.Event{
			Title:       event["title"],
			PlaceName:   sql.NullString{String: event["place_name"], Valid: true},
			Description: event["description"],
			BuyLink:     sql.NullString{String: (event["buy_link"]), Valid: true},
			HappenAt:    event["happen_at"],
			CityId:      city.Id,
			ArtistId:    artist.Id,
		}
		_, err = seed.services.EventManager.Create(eventEntity)

		if err != nil {
			return err
		}
	}
	return nil
}

func (seed *Seed) getEvents() []map[string]string {
	// TODO load from source
	return []map[string]string{
		{
			"artist_name": "Каста",
			"title":       "Каста в Казани c новым альбомом",
			"place_name":  "Максимиллианс",
			"description": "Самое время забыть о дедлайнах, совещаниях и прочих важных делах. Научитесь у рэперов главному — свободе и умению наслаждаться каждым моментом. ",
			"buy_link":    "https://kazan.maximilians.ru/kasta/",
			"city_name":   "Казань",
			"happen_at":   "2022-03-01 20:00:00.000000 +00:00",
		},
		{
			"artist_name": "Каста",
			"title":       "Каста в Москве c новым альбомом",
			"place_name":  "Adreanalin staduin",
			"description": "Самое время забыть о дедлайнах, совещаниях и прочих важных делах. Научитесь у рэперов главному — свободе и умению наслаждаться каждым моментом.  ",
			"buy_link":    "https://afisha.yandex.ru/artist/kasta?city=moscow",
			"city_name":   "Москва",
			"happen_at":   "2022-03-01 20:00:00.000000 +00:00",
		},
		{
			"artist_name": "Би - 2",
			"title":       "Би-2 с лучшими хитами",
			"place_name":  "Татнефть арена",
			"description": "Захватывающая дух сценография, объёмный 4.1 звук и только любимые песни — от новинок из грядущего альбома до ваших любимых хитов!",
			"buy_link":    "https://bdva.ru/concerts/sbp_november_vid_28/",
			"city_name":   "Казань",
			"happen_at":   "2022-03-20 20:00:00.000000 +00:00",
		},
		{
			"artist_name": "Порнофильмы",
			"title":       "Порнофильмы c новым альбомом",
			"place_name":  "Big twin арена",
			"description": "Группа «Порнофильмы» презентует новый альбом в Казани",
			"buy_link":    "https://afisha.yandex.ru/kazan/concert/plan-lomonosova-2022-02-17",
			"city_name":   "Казань",
			"happen_at":   "2022-03-22 19:00:00.000000 +00:00",
		},
	}
}
