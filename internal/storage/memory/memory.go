package memory

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mpuzanov/otus-go/calendar/internal/errors"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStore структура хранения списка событий
type EventStore struct {
	db map[string]model.Event
}

//NewEventStore Возвращаем новое хранилище
func NewEventStore() *EventStore {
	return &EventStore{db: make(map[string]model.Event)}
}

//FindEventByID Найти событие
func (s *EventStore) FindEventByID(id string) (*model.Event, error) {

	ev, exist := s.db[id]
	if exist {
		return &ev, nil
	}
	return nil, errors.ErrNotEvent
}

//AddEvent Добавить событие
func (s *EventStore) AddEvent(event *model.Event) (string, error) {
	event.ID = uuid.New()
	id := event.ID.String()
	s.db[id] = *event
	return id, nil
}

//UpdateEvent Изменить событие
func (s *EventStore) UpdateEvent(event *model.Event) (bool, error) {
	id := event.ID.String()
	_, exist := s.db[id]
	if !exist {
		return false, fmt.Errorf("нет события: %s(%s). %w", event.Header, event.ID, errors.ErrEditEvent)
	}
	s.db[id] = *event
	return true, nil
}

//DelEvent Удалить событие
func (s *EventStore) DelEvent(id string) (bool, error) {
	_, exist := s.db[id]
	if !exist {
		return false, fmt.Errorf("нет события с кодом: %s. %w", id, errors.ErrDelEvent)
	}
	delete(s.db, id)
	return true, nil
}

//GetUserEvents Листинг событий пользователя
func (s *EventStore) GetUserEvents(user string) ([]model.Event, error) {
	var out []model.Event
	for _, val := range s.db {
		if val.UserName == user {
			out = append(out, val)
		}
	}
	return out, nil
}

//String Печать списка событий
func (s *EventStore) String() {
	for _, ev := range s.db {
		fmt.Println(ev)
	}
}
