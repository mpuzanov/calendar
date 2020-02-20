package calendar

import (
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/interfaces"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//Calendar сервис календаря
type Calendar struct {
	Events interfaces.EventStorage
}

// NewCalendar - конструктор календаря
func NewCalendar(repository interfaces.EventStorage) *Calendar {
	return &Calendar{Events: repository}
}

//CreateEvent Создаём событие
func (c *Calendar) CreateEvent(user, header, text string, startTime time.Time, endTime time.Time) (*model.Event, error) {
	return c.Events.CreateEvent(user, header, text, startTime, endTime)
}

//AddEvent Добавление события
func (c *Calendar) AddEvent(event *model.Event) error {
	return c.Events.AddEvent(event)
}

//UpdateEvent Редактирование события
func (c *Calendar) UpdateEvent(event *model.Event) error {
	return c.Events.UpdateEvent(event)
}

//DelEvent Удаление события
func (c *Calendar) DelEvent(event *model.Event) error {
	return c.Events.DelEvent(event)
}

//FindEventByID Поиск события
func (c *Calendar) FindEventByID(id string) (*model.Event, error) {
	return c.Events.FindEventByID(id)
}

//GetEvents Выдаём список событий
func (c *Calendar) GetEvents() []model.Event {
	return c.Events.GetEvents()
}
