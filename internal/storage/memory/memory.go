package memory

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStore структура хранения списка событий
type EventStore struct {
	Events []model.Event
}

//NewEventStore Возвращаем новое хранилище
func NewEventStore() *EventStore {
	return &EventStore{Events: make([]model.Event, 0)}
}

//GetEvents Листинг событий
func (s *EventStore) GetEvents() []model.Event {
	return s.Events
}

//FindEventByHeader Найти событие
func (s *EventStore) FindEventByHeader(header string) (*model.Event, error) {
	for _, event := range s.Events {
		if event.Header == header {
			return &event, nil
		}
	}
	return nil, model.ErrNotEvent
}

//AddEvent Добавить событие
func (s *EventStore) AddEvent(event *model.Event) error {
	ev, _ := s.FindEventByHeader(event.Header)
	if ev != nil {
		return fmt.Errorf("событие с заголовком: %s уже существует. %w", event.Header, model.ErrAddEvent)
	}
	s.Events = append(s.Events, *event)

	return nil
}

//SetEvent Изменить событие
func (s *EventStore) SetEvent(event *model.Event) error {
	for i, ev := range s.Events {
		if ev.UUID == event.UUID {
			s.Events[i] = *event
			return nil
		}
	}
	return fmt.Errorf("нет события с заголовком: %s. %w", event.Header, model.ErrEditEvent)
}

//DelEvent Удалить событие
func (s *EventStore) DelEvent(event *model.Event) error {
	for i, ev := range s.Events {
		if ev.UUID == event.UUID && ev.Header == event.Header && ev.Date == event.Date {
			s.Events = append(s.Events[:i], s.Events[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("нет события с кодом: %s. %w", event.UUID, model.ErrDelEvent)
}

//String Печать списка событий
func (s *EventStore) String() {
	for _, ev := range s.Events {
		fmt.Println(ev)
	}
}

//NewEvent Формируем новое событие
func NewEvent(header string, data time.Time) *model.Event {
	return &model.Event{
		UUID:   GenerateUUID(),
		Header: header,
		Date:   data,
	}
}

//GenerateUUID Генерация уникального кода
func GenerateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
