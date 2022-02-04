package services

import (
	"time"

	"github.com/ainurqa95/expect-artist/pkg/entities"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
)

const comingEventsMaxMonths = 6

type EventService struct {
	repository repositories.EventRepository
}

func NewEventService(repository repositories.EventRepository) *EventService {
	return &EventService{repository: repository}
}

func (service *EventService) FindComingEvents(cityId int, artists []entities.Artist) ([]entities.Event, error) {
	city := entities.City{Id: cityId}
	maxHappenAt := time.Now().AddDate(0, comingEventsMaxMonths, 0).Format(entities.HappenAtDateFormat)

	return service.repository.FindEventsBy(city, artists, maxHappenAt)
}

func (service *EventService) PurifyEventsToShow(events []entities.Event) []string {
	eventMessages := []string{}
	for _, event := range events {
		str := service.buildEventMessage(event)
		eventMessages = append(eventMessages, str)
	}

	return eventMessages
}

func (service *EventService) buildEventMessage(event entities.Event) string {
	str := ""
	str = str + event.Title + "\n"
	if event.PlaceName.Valid {
		str = str + event.PlaceName.String + "\n"
	}
	str = str + event.Description + "\n"

	happenAt, _ := time.Parse(time.RFC3339, event.HappenAt)
	str = str + happenAt.Format("2006-01-02 15:04") + "\n"
	str = str + event.BuyLink.String + "\n"

	return str
}

func (service *EventService) Create(event entities.Event) (int, error) {
	return service.repository.Create(event)
}
