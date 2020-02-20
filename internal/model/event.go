package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//Event структура хранения события
type Event struct {
	//уникальный идентификатор события
	ID uuid.UUID
	// заголовок
	Header string
	// Описание события
	Text string
	// Дата и время события
	StartTime time.Time
	// Дата окончания события
	EndTime time.Time
	// Пользователь, владелец события
	User string
	// За сколько времени высылать уведомление
	ReminderBefore time.Duration
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
