package memslice

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mpuzanov/calendar/internal/errors"
	"github.com/mpuzanov/calendar/internal/model"
)

//EventStore структура хранения списка событий
type EventStore struct {
	db []model.Event
}

//NewEventStore Возвращаем новое хранилище
func NewEventStore() *EventStore {
	return &EventStore{db: make([]model.Event, 0)}
}

//AddEvent Добавить событие
func (s *EventStore) AddEvent(event *model.Event) (string, error) {
	event.ID = uuid.New()
	s.db = append(s.db, *event)
	return event.ID.String(), nil
}

//FindEventByID Найти событие
func (s *EventStore) FindEventByID(id string) (*model.Event, error) {
	for _, event := range s.db {
		if event.ID.String() == id {
			return &event, nil
		}
	}
	return nil, errors.ErrNotEvent
}

//UpdateEvent Изменить событие
func (s *EventStore) UpdateEvent(event *model.Event) (bool, error) {
	for i, ev := range s.db {
		if ev.ID == event.ID {
			s.db[i] = *event
			return true, nil
		}
	}
	return false, fmt.Errorf("нет события: %s(%s). %w", event.Header, event.ID, errors.ErrEditEvent)
}

//DelEvent Удалить событие
func (s *EventStore) DelEvent(id string) (bool, error) {
	iduuid, err := uuid.Parse(id)
	if err != nil {
		return false, err
	}
	for i, ev := range s.db {
		if ev.ID == iduuid {
			s.db = append(s.db[:i], s.db[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("нет события с кодом: %s. %w", iduuid, errors.ErrDelEvent)
}

//String Печать списка событий
func (s *EventStore) String() {
	for _, ev := range s.db {
		fmt.Println(ev)
	}
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
