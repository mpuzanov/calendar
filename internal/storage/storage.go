package storage

import (
	"context"

	"github.com/mpuzanov/calendar/internal/config"
	"github.com/mpuzanov/calendar/internal/interfaces"
	"github.com/mpuzanov/calendar/internal/storage/memory"
	"github.com/mpuzanov/calendar/internal/storage/memslice"
	"github.com/mpuzanov/calendar/internal/storage/postgresdb"
)

//NewStorageDB create storage for calendar
func NewStorageDB(cfg *config.Config) (*interfaces.EventStorage, error) {
	var err error
	var db interfaces.EventStorage

	switch cfg.DB.DbName {
	case "MemorySlice": // MemorySlice хранение событий в slice
		db = memslice.NewEventStore()
	case "MemoryMap": // MemoryMap хранение событий в map
		db = memory.NewEventStore()

	case "Postgres": //Postgres работа с БД
		ctx := context.Background()
		db, err = postgresdb.NewPgEventStore(ctx, cfg.DB.DatabaseURL)
		if err != nil {
			return nil, err
		}
	}
	return &db, err
}
