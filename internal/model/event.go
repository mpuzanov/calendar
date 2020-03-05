package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//Event структура хранения события
type Event struct {
	//уникальный идентификатор события
	ID uuid.UUID `db:"id"`
	// заголовок
	Header string `db:"header"`
	// Описание события
	Text string `db:"text"`
	// Дата и время события
	StartTime time.Time `db:"start_time"`
	// Дата окончания события
	EndTime time.Time `db:"end_time"`
	// Пользователь, владелец события
	UserName string `db:"user_id"`
	// За сколько времени высылать уведомление
	ReminderBefore time.Duration `db:"reminder_before"`
}

//String Строковое представление события
func (e Event) String() string {
	return fmt.Sprintf("ID: %s, событие: %s, начало: %s, окончание: %s\n",
		e.ID,
		e.Header,
		e.StartTime.Format("2006.01.02 15.04.05"),
		e.EndTime.Format("2006.01.02 15.04.05"),
	)
}
