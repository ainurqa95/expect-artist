package config

import (
	"os"

	"github.com/spf13/viper"
)

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start          string `mapstructure:"start"`
	UnknownCommand string `mapstructure:"unknown_command"`
}

type Errors struct {
	Default string `mapstructure:"default"`
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

	// if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
	// 	return err
	// }

	// if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
	// 	return err
	// }

	return nil
}

func fromEnv(cfg *Config) error {
	os.Setenv("DB_NAME", "expect_artist")
	os.Setenv("DB_USERNAME", "default")
	os.Setenv("DB_PASSWORD", "secret")
	// os.Setenv("DB_HOST", "db-expect-artist")
	os.Setenv("TOKEN", "5098762490:AAGNP4Ln0BLhelLfhSSgAe254ZaIWKimuaM")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")

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
