package grpc

import (
	"github.com/mpuzanov/calendar/internal/calendar"
	"github.com/mpuzanov/calendar/internal/config"
	"github.com/mpuzanov/calendar/internal/grpcserver"
	"github.com/mpuzanov/calendar/internal/storage"
	"github.com/mpuzanov/calendar/pkg/logger"

	"log"

	"github.com/spf13/cobra"
)

var cfgPath string

var (
	// ServerCmd .
	ServerCmd = &cobra.Command{
		Use:   "grpc_server",
		Short: "Run grpc server",
		Run:   grpcServerStart,
	}
)

func init() {
	ServerCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
}

func grpcServerStart(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("unable to load %s: %s", cfgPath, err)
	}
	logger := logger.NewLogger(cfg.Log)
	// Init db
	db, err := storage.NewStorageDB(cfg)
	if err != nil {
		log.Fatalf("newStorageDB failed: %s", err)
	}
	calendar := calendar.NewCalendar(db)

	if err := grpcserver.Start(cfg, logger, calendar); err != nil {
		log.Fatal(err)
	}
}
