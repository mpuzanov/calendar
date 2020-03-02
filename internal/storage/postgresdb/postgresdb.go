package postgresdb

import (
	"context"
	"database/sql"

	// Register some standard stuff
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mpuzanov/otus-go/calendar/internal/errors"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//EventStore структура хранения списка событий Календаря
type EventStore struct {
	ctx context.Context
	db  *sqlx.DB
}

//NewPgEventStore Возвращаем хранилище
func NewPgEventStore(ctx context.Context, databaseURL string) (*EventStore, error) {
	db, err := sqlx.Connect("pgx", databaseURL)
	if err != nil {
		return nil, err
	}
	return &EventStore{ctx: ctx, db: db}, nil
}

//AddEvent Добавляем событие
func (pg *EventStore) AddEvent(event *model.Event) (string, error) {
	eDb, err := toEventDB(event)
	if err != nil {
		return "", err
	}
	query1 := "INSERT INTO events (header, text, start_time, end_time, user_name, reminder_before) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var id string
	if err := pg.db.QueryRowContext(pg.ctx, query1,
		eDb.Header, eDb.Text, eDb.StartTime, eDb.EndTime, eDb.UserName, eDb.ReminderBefore.Get()).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

//UpdateEvent Изменение события
func (pg *EventStore) UpdateEvent(event *model.Event) (bool, error) {
	eDb, err := toEventDB(event)
	if err != nil {
		return false, err
	}
	query := `UPDATE events SET (header, text, start_time, end_time, user_name, reminder_before) = 
	(:Header, :Text, :StartTime, :EndTime, :UserName, :ReminderBefore) WHERE ID = :ID`
	result, err := pg.db.NamedExecContext(pg.ctx, query,
		map[string]interface{}{
			"ID":             eDb.ID,
			"Header":         eDb.Header,
			"Text":           eDb.Text,
			"StartTime":      eDb.StartTime,
			"EndTime":        eDb.EndTime,
			"UserName":       eDb.UserName,
			"ReminderBefore": eDb.ReminderBefore.Get(),
		})

	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.ErrNoDBAffected
	}
	return true, nil
}

//DelEvent Удаляем событие по коду
func (pg *EventStore) DelEvent(id string) (bool, error) {
	res, err := pg.db.ExecContext(pg.ctx, "DELETE FROM events WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.ErrNoDBAffected
	}
	return true, nil
}

//FindEventByID возврат события по коду
func (pg *EventStore) FindEventByID(id string) (*model.Event, error) {
	evDB := &EventDB{}
	query := "SELECT * FROM events WHERE id= $1"
	if err := pg.db.GetContext(pg.ctx, evDB, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}
	ev, err := toEvent(evDB)
	if err != nil {
		return nil, err
	}
	return ev, nil
}

//GetUserEvents получаем события пользователя
func (pg *EventStore) GetUserEvents(user string) ([]model.Event, error) {
	query := `select * from events where user_name=$1`

	out := make([]model.Event, 0)
	outDB := make([]EventDB, 0)
	err := pg.db.Select(&outDB, query, user)
	if err == sql.ErrNoRows {
		return out, nil
	}
	if err != nil {
		return nil, err
	}
	for _, v := range outDB {
		e, err := toEvent(&v)
		if err != nil {
			return nil, err
		}
		out = append(out, *e)
	}
	return out, nil
}
