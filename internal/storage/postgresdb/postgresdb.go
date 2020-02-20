package postgresdb

import (
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStore implements Calendar
type EventStore struct {
	db *sqlx.DB
}

//NewPgEventStore ...
func NewPgEventStore(dsn string) (*EventStore, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &EventStore{db: db}, nil
}

//CreateEvent ...
func (pg *EventStore) CreateEvent(user, header, text string, startTime time.Time, endTime time.Time) (*model.Event, error) {
	var ev model.Event
	return &ev, nil
}

//AddEvent ...
func (pg *EventStore) AddEvent(event *model.Event) error {
	return nil
}

//UpdateEvent ...
func (pg *EventStore) UpdateEvent(event *model.Event) error {
	return nil
}

//DelEvent ...
func (pg *EventStore) DelEvent(event *model.Event) error {
	return nil
}

//FindEventByID ...
func (pg *EventStore) FindEventByID(id string) (*model.Event, error) {
	var ev model.Event
	return &ev, nil
}

//GetEvents ...
func (pg *EventStore) GetEvents() []model.Event {
	var out []model.Event
	return out
}
