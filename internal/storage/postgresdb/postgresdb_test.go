package postgresdb_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mpuzanov/calendar/internal/model"
	"github.com/mpuzanov/calendar/internal/storage/postgresdb"
	"github.com/stretchr/testify/assert"
)

var (
	dbURL     string
	UserTest  string = "User1"
	eventTest *model.Event
)

func TestMain(m *testing.M) {
	dbURL = os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "host=localhost dbname=pg_calendar_test user=postgres password=12345 sslmode=disable"
		//databaseURL = "postgres://postgres:12345@localhost:5432/pg_calendar?sslmode=disable"
	}

	startTime := time.Date(2020, time.April, 1, 9, 0, 0, 0, time.UTC)
	endTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 10:30")
	reminderBefore := time.Duration(time.Minute * 30)
	textEvent := "описание события"
	eventTest = &model.Event{Header: "Событие 1", Text: textEvent, StartTime: startTime, EndTime: endTime, UserName: UserTest, ReminderBefore: reminderBefore}

	os.Exit(m.Run())
}

func TestAddEvent(t *testing.T) {
	ctx := context.Background()
	pg, _ := postgresdb.NewPgEventStore(ctx, dbURL)
	got, err := pg.AddEvent(eventTest)
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
}

func TestGetUserEvents(t *testing.T) {
	ctx := context.Background()
	pg, _ := postgresdb.NewPgEventStore(ctx, dbURL)
	got, err := pg.GetUserEvents(UserTest)
	assert.NoError(t, err)
	assert.NotNil(t, got)
}

func TestFindEventByID(t *testing.T) {
	ctx := context.Background()
	pg, _ := postgresdb.NewPgEventStore(ctx, dbURL)

	evList, err := pg.GetUserEvents(UserTest)
	assert.NoError(t, err)
	//t.Logf("Кол-во записей: %v", len(evList))
	if len(evList) == 0 {
		t.Log("Таблица событий пуста")
		t.Skip()
	}
	id := evList[0].ID.String()
	//t.Logf("Ищем %s", id)
	got, err := pg.FindEventByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, got)
}

func TestUpdateEvent(t *testing.T) {
	ctx := context.Background()
	pg, _ := postgresdb.NewPgEventStore(ctx, dbURL)

	//добавляем событие
	id, err := pg.AddEvent(eventTest)
	assert.NoError(t, err)
	assert.NotEqual(t, "", id)
	eventTest.ID = uuid.MustParse(id)
	newText := "Новое описание события"
	eventTest.Text = newText
	//изменяем событие
	got, err := pg.UpdateEvent(eventTest)
	assert.NoError(t, err)
	assert.True(t, got)
}

func TestDelEvent(t *testing.T) {
	ctx := context.Background()
	pg, _ := postgresdb.NewPgEventStore(ctx, dbURL)

	//добавляем событие
	id, err := pg.AddEvent(eventTest)
	assert.NoError(t, err)
	assert.NotEqual(t, "", id)
	//удаляем событие
	got, err := pg.DelEvent(id)
	assert.NoError(t, err)
	assert.True(t, got)
}
