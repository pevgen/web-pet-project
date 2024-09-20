package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type WebServerConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	Postgres PgConfig    `yaml:"postgres"`
	Mongodb  MongoConfig `yaml:"mongodb"`
}

type PgConfig struct {
	ConnectString string `yaml:"connect-string"`
}

type MongoConfig struct {
	ConnectString string `yaml:"connect-string"`
	DbName        string `yaml:"db-name"`
}

type AppConfig struct {
	WebServer WebServerConfig `yaml:"web-server"`
	Db        DBConfig        `yaml:"db"`
}

func New(configPath string) AppConfig {

	log.Printf("Start reading configuration from file: %v\n", configPath)

	var cfg AppConfig

	bytes, err := os.ReadFile(configPath) // Чтение и анализ YAML-файла
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
