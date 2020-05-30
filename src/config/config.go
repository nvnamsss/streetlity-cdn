package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Configuration struct {
	ResourcesLocation string `json:"resources-location"`
}

var Config Configuration

func LoadConfig(path string) {
	file, fileErr := os.Open(path)
	if fileErr != nil {

		log.Panic(fileErr)
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = Configuration{}

	err := decoder.Decode(&Config)

	if err != nil {
		log.Panic(err)
	}
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	configPath := filepath.Join(filepath.Dir(basepath), "config", "config.json")

	fmt.Println("[Config]", configPath)
	LoadConfig(configPath)
}
