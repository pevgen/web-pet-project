package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigurationYaml struct {
	Enabled bool
	Path    string
}

func ReadYaml(configName string) ConfigurationYaml {
	var cfg ConfigurationYaml

	configFile, err := os.ReadFile(configName) // Чтение и анализ YAML-файла
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Path: %v \n", cfg.Path)
	fmt.Printf("enabled: %v \n", cfg.Enabled)
	return cfg
}
