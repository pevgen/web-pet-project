package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type configurationJson struct {
	Enabled bool
	Path    string
}

func ReadJson(configName string) {
	var conf configurationJson

	file, _ := os.Open(configName) // Открытие	конфигурационного	файла
	defer file.Close()
	decoder := json.NewDecoder(file) //Извлечение	JSON - значений	в	переменные
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
	fmt.Println(conf.Enabled)
}
