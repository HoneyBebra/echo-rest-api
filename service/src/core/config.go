package core

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug              string
	ApiPrefix          string
	HTTPAddr           string // You also need to change port in the nginx settings
	ReadTimeoutMinute  int
	WriteTimeoutMinute int
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
		readTimeoutMinute, err := strconv.Atoi(os.Getenv("READ_TIMEOUT_MINUTE"))
		if err != nil {
			log.Fatal(err)
		}
		writeTimeoutMinute, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT_MINUTE"))
		if err != nil {
			log.Fatal(err)
		}

		ConfigInstance = &Config{
			Debug:              os.Getenv("DEBUG"),
			ApiPrefix:          os.Getenv("API_PREFIX"),
			HTTPAddr:           os.Getenv("HTTP_ADDR"),
			ReadTimeoutMinute:  readTimeoutMinute,
			WriteTimeoutMinute: writeTimeoutMinute,
		}
	})
	return ConfigInstance
}
