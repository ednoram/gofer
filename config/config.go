package config

import (
	"log"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
	Client struct {
		ApiUrl        string `toml:"api_url"`
		ApiKeyVarName string `toml:"api_key_var_name"`
	} `toml:"client"`
}

var (
	config   *Config
	loadOnce sync.Once
	filePath = "config/config.toml"
)

func GetConfig() *Config {
	loadOnce.Do(func() {
		if _, err := toml.DecodeFile(filePath, &config); err != nil {
			log.Fatalf("Error decoding TOML file: %v", err)
		}
	})

	return config
}
