package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port       string
	Dir        string
	LogFileDir string
}

func Load() Config {

	file, err := os.Open("config/config")
	if err != nil {
		fmt.Println("Error opening configuration file:", err)
	}
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error decoding configuration file:", err)
	}

	return configuration
}
