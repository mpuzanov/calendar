package grpcserver

import (
	"log"
	"net"

	"github.com/mpuzanov/calendar/internal/calendar"
	"github.com/mpuzanov/calendar/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//Start ...
func Start(conf *config.Config, logger *zap.Logger, evs *calendar.Calendar) error {

	s := &GRPCServer{
		cfg:          conf,
		logger:       logger,
		eventService: evs,
	}

	l, err := net.Listen("tcp", s.cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("Cannot listen: %s\n", err)
	}

	grpcServer := grpc.NewServer()

	RegisterCalendarServer(grpcServer, s)

	log.Printf("Starting gRPC server %s, file log: %s\n", s.cfg.GRPCAddr, s.cfg.Log.LogFile)
	s.logger.Info("Starting gRPC server", zap.String("address", s.cfg.GRPCAddr))

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Cannot start gRPC server: %s\n", err)
	}

	return nil
}
