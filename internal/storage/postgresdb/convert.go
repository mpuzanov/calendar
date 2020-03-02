package postgresdb

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/pgtype"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

const (
	day   int64 = int64(time.Hour) * int64(24)
	month int64 = day * 30
)

func toInterval(t pgtype.Interval) time.Duration {
	d := time.Duration(t.Microseconds*int64(time.Microsecond) + int64(t.Days)*day + int64(t.Months)*month)
	return d
}

func toNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

//toEvent конвертируем структуру базы данных в model.Event
func toEvent(d *EventDB) (*model.Event, error) {
	var Text string
	if d.Text.Valid {
		Text = d.Text.String
	}

	id, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}

	return &model.Event{
		ID:             id,
		Header:         d.Header,
		Text:           Text,
		StartTime:      d.StartTime,
		EndTime:        d.EndTime,
		UserName:       d.UserName,
		ReminderBefore: toInterval(d.ReminderBefore),
	}, nil
}

//toEventDB конвертируем структуру model.Event в аналог базы данных
func toEventDB(ev *model.Event) (*EventDB, error) {
	var ReminderBefore pgtype.Interval
	if err := ReminderBefore.Set(ev.ReminderBefore); err != nil {
		return nil, err
	}

	return &EventDB{
		ID:             ev.ID.String(),
		Header:         ev.Header,
		Text:           toNullString(ev.Text),
		StartTime:      ev.StartTime,
		EndTime:        ev.EndTime,
		UserName:       ev.UserName,
		ReminderBefore: ReminderBefore,
	}, nil
}
