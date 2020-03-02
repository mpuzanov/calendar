package main

import (
	"log"

	"github.com/mpuzanov/otus-go/calendar/internal/calendar"
	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/grpcserver"
	"github.com/mpuzanov/otus-go/calendar/internal/storage"

	//"github.com/mpuzanov/otus-go/calendar/internal/web"

	"github.com/mpuzanov/otus-go/calendar/pkg/logger"
	flag "github.com/spf13/pflag"
)

var configFile string

func init() {
	flag.StringVarP(&configFile, "config", "c", "", "path config file")
	flag.Parse()
}

func main() {

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", configFile, err)
	}

	logger := logger.NewLogger(cfg.Log)

	// Init db
	db, err := storage.NewStorageDB(cfg)
	if err != nil {
		log.Fatalf("newStorageDB failed: %s", err)
	}

	calendar := calendar.NewCalendar(*db)

	//sampleWorkCalendar(calendar)

	// if err := web.Start(cfg, logger, calendar); err != nil {
	// 	log.Fatal(err)
	// }

	if err := grpcserver.Start(cfg, logger, calendar); err != nil {
		log.Fatal(err)
	}

}

// ./calendar --config=configs/config-dev.yml
// ./calendar --config=configs/config-prod.yml
// grpc_listen=":50052" ./calendar --config=configs/config-dev.yml
// http_listen=":8091" ./calendar --config=configs/config-dev.yml
// curl -i localhost:8091
// curl -i localhost:8091/hello/Mikhail
