package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/mpuzanov/calendar/internal/calendar"
	"github.com/mpuzanov/calendar/internal/config"
	"go.uber.org/zap"
)

// Start Запуск сервера
func Start(conf *config.Config, logger *zap.Logger, evs *calendar.Calendar) error {

	srv := newServer(conf, logger, evs)

	//запускаем веб-сервер
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server http started: %s, file log: %s\n", srv.Addr, conf.Log.LogFile)
	logger.Info("Starting Http server", zap.String("address", srv.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Print("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
	log.Println("Shutdown done")
	os.Exit(0)

	return nil
}

func newServer(config *config.Config, logger *zap.Logger, evs *calendar.Calendar) *http.Server {

	handler := &myHandler{
		router:       mux.NewRouter(),
		logger:       logger,
		eventService: evs,
	}
	handler.configRouter()

	server := &http.Server{
		Addr:           config.HTTPAddr,
		Handler:        handler,
		IdleTimeout:    10 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server
}
