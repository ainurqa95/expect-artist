package seeders

import (
	"fmt"

	"github.com/ainurqa95/expect-artist/pkg/entities"
)

func (seed *Seed) SeedMessageTypes() error {
	existsData := seed.repositories.MessageTypeRepository.ExistsData()
	fmt.Println(existsData)
	if existsData {
		return nil
	}

	return seed.saveMessageTypes()
}

func (seed *Seed) saveMessageTypes() error {
	messageTypes := seed.getMessageTypes()
	for _, messageType := range messageTypes {
		entity := entities.MessageType{
			Name: messageType["name"],
			Code: messageType["code"],
		}
		_, err := seed.repositories.MessageTypeRepository.Create(entity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (seed *Seed) getMessageTypes() []map[string]string {
	return []map[string]string{
		{"code": entities.SearchArtistCommand, "name": "Команда /search"},
		{"code": entities.AfterSearchArtistCommand, "name": "Сообщение после нажатие на search"},
		{"code": entities.ChosedArtist, "name": "Сообщение выбран артист из списка или отклонен"},
		{"code": entities.SetUpCityCommand, "name": "Команда /set_up_city"},
		{"code": entities.AfterSetUpCityCommand, "name": "Сообщение после нажатия /set_up_city"},
		{"code": entities.ChosedCity, "name": "Сообщение выбран город"},
		{"code": entities.OtherMessageCommand, "name": "Другое сообщение или команда"},
		{"code": entities.EventsCommand, "name": "Команда /events"},
	}

}
