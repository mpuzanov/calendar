package main

import (
	"log"

	flag "github.com/spf13/pflag"

	"github.com/mpuzanov/otus-go/calendar/internal/apiserver"
)

var configFile string

func init() {
	flag.StringVarP(&configFile, "config", "c", "configs/config.yml", "path config file")
	flag.Parse()
}

func main() {
	var err error

	cfg, err := apiserver.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", configFile, err)
	}

	if err = apiserver.Start(cfg); err != nil {
		log.Fatal(err)
	}

}

// ./calendar --config=configs/config-dev.yml
// PORT=8091 ./calendar --config=configs/config-dev.yml
// ./calendar --config=configs/config-prod.yml
