package main

import (
	"fmt"
	"log"

	"github.com/ainurqa95/expect-artist/pkg/config"
	"github.com/ainurqa95/expect-artist/pkg/repositories"
	"github.com/ainurqa95/expect-artist/pkg/seeders"
	"github.com/ainurqa95/expect-artist/pkg/services"
	"github.com/ainurqa95/expect-artist/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // <------------ here
)

func main() {
	config, err := config.Init()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.Db.Name)
	fmt.Println(config.Db.Host)
	fmt.Println(config.Db.Port)
	fmt.Println(config.Db.UserName)
	fmt.Println(config.Db.Password)
	fmt.Println(config.BotURL)
	fmt.Println(config.BotToken)
	db, err := NewPostgresDB(config.Db)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repositories.NewRepository(db)
	services := services.NewService(repos)
	seeder := seeders.NewSeed(repos)
	err = seeder.SeedData()
	bot := InitTelegramBot(config, services)
	bot.Start()
}

func InitTelegramBot(config *config.Config, services *services.Service) *telegram.Bot {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	return telegram.NewBot(bot, services, config.Messages)
}

func NewPostgresDB(cfg config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.UserName, cfg.Name, cfg.Password, "disable"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
