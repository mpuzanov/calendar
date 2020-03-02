package web

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mpuzanov/otus-go/calendar/internal/calendar"
	"go.uber.org/zap"
)

type myHandler struct {
	router       *mux.Router
	logger       *zap.Logger
	eventService *calendar.Calendar
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
