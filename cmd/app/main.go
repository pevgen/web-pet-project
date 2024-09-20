package main

import (
	"os"
	"web-pet-project/internal/config"
	"web-pet-project/internal/web"
)

func main() {

	// Config can be set with an environment variable CONFIG_PATH
	configPath, exists := os.LookupEnv("CONFIG_PATH")

	// If the environment variable CONFIG_PATH doesn't exist
	// then default path is "config/app-config.yaml"
	if !exists {
		configPath = "config/app-config.yaml"
	}

	c := config.New(configPath)

	web.NewRouter(c)

}
