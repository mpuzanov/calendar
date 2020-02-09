package model

import (
	"fmt"
	"time"
)

//Event структура хранения события
type Event struct {
	//уникальный идентификатор события
	UUID string
	// заголовок
	Header string
	// Дата и время события
	Date time.Time
	// Длительность события
	TimeDuration time.Duration
	// Описание события
	Description string
	// Пользователь, владелец события
	User
}

//String Строковое представление события
func (e Event) String() string {
	return fmt.Sprintf("заголовок: %s, UUID: %s, Время: %s\n", e.Header, e.UUID, e.Date.Format("2006.01.02 15.04.05"))
}
