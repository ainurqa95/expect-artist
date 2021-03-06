package config

import (
	"github.com/spf13/viper"
)

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start                        string `mapstructure:"start"`
	Help                         string `mapstructure:"start"`
	Search                       string `mapstructure:"search"`
	ShowEvents                   string `mapstructure:"show_events"`
	SetUpCity                    string `mapstructure:"set_up_city"`
	SelectArtist                 string `mapstructure:"select_artist"`
	SelectCity                   string `mapstructure:"select_city"`
	UnknownCommand               string `mapstructure:"unknown_command"`
	SubscriptionAlreadyExists    string `mapstructure:"subscription_already_exists"`
	SubscriptionSuccesfullyAdded string `mapstructure:"subsscription_succesfully_added"`
	CitySettingUpIsSuccesfull    string `mapstructure:"city_setting_up_is_succesfull"`
}

type Errors struct {
	Default               string `mapstructure:"default"`
	ArtistNotFound        string `mapstructure:"unknown_artist"`
	CityNotFound          string `mapstructure:"unknown_city"`
	CityHasNotBeenSet     string `mapstructure:"city_has_not_been_set"`
	ArtistsHaveNotBeenSet string `mapstructure:"artists_have_not_been_set"`
	EventsNotFound        string `mapstructure:"events_not_found"`
}

type DB struct {
	Host     string
	Name     string
	Port     string
	UserName string
	Password string
}

type Config struct {
	Db       DB
	BotToken string
	BotURL   string `mapstructure:"bot_url"`
	Messages Messages
}

func Init() (*Config, error) {

	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {

	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	cfg.BotToken = viper.GetString("token")

	if err := viper.BindEnv("db_name"); err != nil {
		return err
	}
	cfg.Db.Name = viper.GetString("db_name")

	if err := viper.BindEnv("DB_USERNAME"); err != nil {
		return err
	}
	cfg.Db.UserName = viper.GetString("DB_USERNAME")

	if err := viper.BindEnv("db_password"); err != nil {
		return err
	}
	cfg.Db.Password = viper.GetString("db_password")

	if err := viper.BindEnv("db_port"); err != nil {
		return err
	}
	cfg.Db.Port = viper.GetString("db_port")

	if err := viper.BindEnv("db_host"); err != nil {
		return err
	}
	cfg.Db.Host = viper.GetString("db_host")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
