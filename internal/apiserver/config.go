package apiserver

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config Структура файла с конфигурацией
type Config struct {
	BindAddr      string `yaml:"http_listen"`
	LogLevel      string `yaml:"log_level"`
	LogFile       string `yaml:"log_file"`
	LogFormatJSON bool   `yaml:"log_format_JSON"`
	DatabaseURL   string `yaml:"database_url"`
}

// LoadConfig Загрузка конфигурации из файла
func LoadConfig(filepath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err = yaml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	port, exists := os.LookupEnv("PORT")
	if exists { //Заменяем порт
		//log.Printf("PORT: %s\n", port)
		var splits = strings.Split(config.BindAddr, ":")
		config.BindAddr = splits[0] + ":" + port
	}

	return config, nil
}
