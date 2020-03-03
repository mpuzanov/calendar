package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

//Config Структура файла с конфигурацией
type Config struct {
	Log      LogConf `yaml:"log" mapstructure:"log"`
	DB       DBConf  `yaml:"db" mapstructure:"db"`
	HTTPAddr string  `yaml:"http_listen" mapstructure:"http_listen"`
	GRPCAddr string  `yaml:"grpc_listen" mapstructure:"grpc_listen"`
}

// LogConf стуктура для настройки логирования
type LogConf struct {
	LogLevel      string `yaml:"loglevel" mapstructure:"loglevel"`
	LogFile       string `yaml:"logfile" mapstructure:"logfile"`
	LogFormatJSON bool   `yaml:"logformat_JSON" mapstructure:"logformat_JSON"`
}

// DBConf стуктура для настройки работы с базой данных
type DBConf struct {
	DbName      string `yaml:"db_name" mapstructure:"db_name"`
	DatabaseURL string `yaml:"database_url" mapstructure:"url"`
}

// LoadConfig Загрузка конфигурации из файла
func LoadConfig(filePath string) (*Config, error) {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("log.loglevel", "info")
	viper.SetDefault("http_listen", "0.0.0.0:8090")
	viper.SetDefault("grpc_listen", "0.0.0.0:50051")
	viper.SetDefault("db.url", "postgres://postgres:12345@localhost:5432/pg_calendar_test?sslmode=disable")

	if filePath != "" {
		log.Printf("Parsing config: %s\n", filePath)
		viper.SetConfigFile(filePath)
		viper.SetConfigType("yaml")
		//log.Println(viper.ConfigFileUsed())
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("Config file is not specified.")
	}
	//log.Println(viper.AllSettings())

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	//log.Println(config)
	return &config, nil
}
