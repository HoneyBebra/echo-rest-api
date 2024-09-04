package core

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug     string
	ApiPrefix string
	Port      string // You also need to change it in the nginx settings
}

// Get config once
var (
	ConfigInstance *Config
	once           sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		ConfigInstance = &Config{
			Debug:     os.Getenv("DEBUG"),
			ApiPrefix: os.Getenv("API_PREFIX"),
			Port:      os.Getenv("PORT"),
		}
	})
	return ConfigInstance
}
