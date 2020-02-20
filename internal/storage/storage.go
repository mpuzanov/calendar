package storage

import (
	"github.com/mpuzanov/otus-go/calendar/internal/interfaces"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memory"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/memslice"
	"github.com/mpuzanov/otus-go/calendar/internal/storage/postgresdb"
)

//NewStorageDB create storage for calendar
func NewStorageDB(dbname, dburl string) (*interfaces.EventStorage, error) {
	var err error
	var db interfaces.EventStorage

	switch dbname {
	case "MemorySlice": // MemorySlice хранение событий в slice
		db = memslice.NewEventStore()
	case "MemoryMap": // MemoryMap хранение событий в map
		db = memory.NewEventStore()

	case "Postgres": //Postgres работа с БД
		db, err = postgresdb.NewPgEventStore(dburl)
		if err != nil {
			return nil, err
		}
	}
	return &db, err
}
