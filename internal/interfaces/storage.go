package interfaces

import (
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStorage интерфейс для работы с DB
type EventStorage interface {
	CreateEvent(user, header, text string, startTime time.Time, endTime time.Time) (*model.Event, error)
	AddEvent(event *model.Event) error
	UpdateEvent(event *model.Event) error
	DelEvent(event *model.Event) error
	FindEventByID(id string) (*model.Event, error)
	GetEvents() []model.Event
}
