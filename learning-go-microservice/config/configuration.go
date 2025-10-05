package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Server struct {
	BasePath string
	Port     string
}

type Config struct {
	Server Server
}

func LoadConfig() (Config, error) {
	var config Config

	data, err := os.ReadFile("/Users/aniruddhasharma/Desktop/Folder/Projects/robust-go-microservice/learning-go-microservice/exe/application.yaml") // Use os.ReadFile instead of ioutil.ReadFile
	fmt.Println("Data: " + string(data))
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
