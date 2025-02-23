package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
	Client struct {
		ApiUrl string `toml:"api_url"`
	} `toml:"client"`
}

var (
	config   *Config
	loadOnce sync.Once
)

func GetConfig() *Config {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	filePath := filepath.Join(cwd, "config", "config.toml")

	loadOnce.Do(func() {
		if _, err := toml.DecodeFile(filePath, &config); err != nil {
			log.Fatalf("Error decoding TOML file: %v", err)
		}
	})

	return config
}
