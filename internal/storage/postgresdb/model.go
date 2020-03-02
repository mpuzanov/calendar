package postgresdb

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/pgtype"
)

// EventDB - структура события для postgres
type EventDB struct {
	//уникальный идентификатор события
	//ID pgtype.UUID `db:"id"`
	ID string `db:"id"`
	// заголовок
	Header string `db:"header"`
	// Описание события
	Text sql.NullString `db:"text"`
	// Дата и время события
	StartTime time.Time `db:"start_time"`
	// Дата окончания события
	EndTime time.Time `db:"end_time"`
	// Пользователь, владелец события
	UserName string `db:"user_name"`
	// За сколько времени высылать уведомление
	ReminderBefore pgtype.Interval `db:"reminder_before"`
}
