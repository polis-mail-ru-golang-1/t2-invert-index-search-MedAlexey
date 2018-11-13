package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Listen string
	Dir    string
}

var Configuration Config

func Load(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening configuration file:", err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	if err != nil {
		fmt.Println("Error decoding configuration file:", err)
		os.Exit(1)
	}

}
