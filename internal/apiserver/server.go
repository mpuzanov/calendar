package apiserver

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/mpuzanov/otus-go/calendar/internal/storage"
	"github.com/mpuzanov/otus-go/calendar/pkg/logger"
	"go.uber.org/zap"
)

type myHandler struct {
	router *mux.Router
	logger *zap.Logger
	store  storage.Storage
}

// Start Запуск сервера
func Start(config *Config) error {

	err := storage.NewStorage(storage.Memory)
	if err != nil {
		return err
	}
	srv := newServer(storage.DB, config)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server started: ", srv.Addr)

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

func newServer(store storage.Storage, config *Config) *http.Server {

	cfglog := logger.Configuration{
		Level:      config.LogLevel,
		JSONFormat: config.LogFormatJSON,
		LogFile:    config.LogFile,
	}

	handler := &myHandler{
		router: mux.NewRouter(),
		logger: logger.NewLogger(cfglog), //logger.InitLogger(cfglog), //
		store:  store,
	}
	handler.configRouter()

	server := &http.Server{
		Addr:           config.BindAddr,
		Handler:        handler,
		IdleTimeout:    10 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server
}

func (s *myHandler) configRouter() {
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/", s.homePage)
	s.router.HandleFunc("/hello", s.helloPage)
	s.router.HandleFunc("/hello/{name}", s.helloPage)
}

func (s *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *myHandler) homePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body><h1>Сервис \"Календарь\"</h1></body></html>")
}

func (s *myHandler) helloPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		name = "незнакомец"
	}
	s.logger.Debug("helloPage",
		zap.String("name", name))
	io.WriteString(w, "<html><body><h1>Добро пожаловать "+name+"!</h1></body></html>")
}

func (s *myHandler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//logger:=s.logger.With(zap.String("remote_addr", r.RemoteAddr))
		s.logger.Info("Request",
			zap.String("Method", r.Method),
			zap.String("URI", r.RequestURI),
			zap.String("remote_addr", r.RemoteAddr),
		)
		next.ServeHTTP(w, r)
	})
}
