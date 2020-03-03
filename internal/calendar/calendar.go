package calendar

import (
	"github.com/mpuzanov/calendar/internal/interfaces"
	"github.com/mpuzanov/calendar/internal/model"
)

//Calendar сервис календаря
type Calendar struct {
	Events interfaces.EventStorage
}

// NewCalendar - конструктор календаря
func NewCalendar(repository interfaces.EventStorage) *Calendar {
	return &Calendar{Events: repository}
}

//AddEvent Добавление события
func (c *Calendar) AddEvent(event *model.Event) (string, error) {
	return c.Events.AddEvent(event)
}

//UpdateEvent Редактирование события
func (c *Calendar) UpdateEvent(event *model.Event) (bool, error) {
	return c.Events.UpdateEvent(event)
}

//DelEvent Удаление события
func (c *Calendar) DelEvent(id string) (bool, error) {
	return c.Events.DelEvent(id)
}

//FindEventByID Поиск события
func (c *Calendar) FindEventByID(id string) (*model.Event, error) {
	return c.Events.FindEventByID(id)
}

//GetUserEvents Выдаём список событий пользователя
func (c *Calendar) GetUserEvents(user string) ([]model.Event, error) {
	return c.Events.GetUserEvents(user)
}
