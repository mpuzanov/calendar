package storage

import (
	"github.com/mpuzanov/otus-go/calendar/internal/model"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memory"
)

const (
	//Postgres работа с БД
	Postgres int = iota
	// Memory будет работа в памяти
	Memory
)

//DB интерфейс для взаимодействия с данными на нескольких уровнях
var DB Storage

//Storage интерфейс для работы со структурой календаря
type Storage interface {
	AddEvent(event *model.Event) error
	SetEvent(event *model.Event) error
	DelEvent(event *model.Event) error
	FindEventByHeader(header string) (*model.Event, error)
	GetEvents() []model.Event
}

//NewStorage создание интерфейс для работы со структурой календаря
func NewStorage(storageType int) error {
	//var err error

	switch storageType {
	case Memory:
		DB = new(memory.EventStore)

		// case Postgres:
		// 	DB, err = storage.InitDB()
		// 	if err != nil {
		// 		return err
		// 	}
	}
	return nil
}
