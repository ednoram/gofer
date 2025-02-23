package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Port   int    `mapstructure:"PORT"`
	ApiUrl string `mapstructure:"GOFER_API_URL"`
	ApiKey string `mapstructure:"GOFER_API_KEY"`
}

var (
	cfg      *Config
	loadOnce sync.Once
)

func initConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
}

func GetConfig() *Config {
	loadOnce.Do(initConfig)

	return cfg
}
