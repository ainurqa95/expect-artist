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
		entity := entities.TelegramMessageType{
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
		{"code": entities.AfterSearchArtistType, "name": "Поиск артиста"},
		{"code": entities.AfterSetUpCityType, "name": "Установка города"},
		{"code": entities.CommadType, "name": "Команда"},
	}

}
