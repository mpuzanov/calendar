package memory

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/mpuzanov/calendar/internal/errors"
	"github.com/mpuzanov/calendar/internal/model"
)

var (
	testStore          *EventStore
	eventTest          model.Event
	startTime, endTime time.Time
	userName           string
)

func TestMain(m *testing.M) {
	var err error
	testStore, err = generateTestData()
	if err != nil {
		log.Fatalf("Ошибка генерации тестовых данных. %v", err)
	}
	// for _, v := range testStore.db {
	// 	fmt.Print(v)
	// }
	os.Exit(m.Run())
}

func generateTestData() (*EventStore, error) {
	userName = "user1"
	startTime = time.Date(2020, time.April, 1, 9, 0, 0, 0, time.UTC)
	endTime, _ = time.Parse("2006-01-02 15:04", "2020-04-01 10:30")
	store := NewEventStore()
	for i := 1; i <= 10; i++ {
		event := &model.Event{Header: "Event " + strconv.Itoa(i),
			Text:           "",
			StartTime:      startTime,
			EndTime:        endTime,
			UserName:       userName,
			ReminderBefore: 0}
		_, err := store.AddEvent(event)
		if err != nil {
			return nil, err
		}
	}
	return store, nil
}

func TestFindEventByID(t *testing.T) {
	var eventFind model.Event
	for _, val := range testStore.db {
		eventFind = val
		break
	}

	testCases := []struct {
		desc string
		find string
		want string
		err  error
	}{
		{
			desc: "Тест 1",
			find: eventFind.ID.String(),
			want: eventFind.ID.String(),
			err:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := testStore.FindEventByID(tC.find)
			if err != nil {
				t.Errorf("%s, got=%v, expected=%v, error: %v", tC.desc, got.ID.String(), tC.want, err)
			}
			if got.ID.String() != tC.want {
				t.Errorf("Error FindEventByHeader. %s, got=%v, expected=%v", tC.desc, got.ID.String(), tC.want)
			}
		})
	}
}

func TestAddEvent(t *testing.T) {
	countEvent := len(testStore.db)
	eventTest = model.Event{UserName: userName, Header: "Event Add", StartTime: startTime, EndTime: endTime}
	testCases := []struct {
		desc        string
		eventHeader string
		err         error
		want        int
	}{
		{
			desc:        "Test (success)",
			eventHeader: "Event Add",
			err:         nil,
			want:        countEvent + 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			if _, err := testStore.AddEvent(&eventTest); err != nil {
				if !errors.Is(err, tC.err) {
					t.Errorf("%s error: %v", tC.desc, err)
				}
			}
			got := len(testStore.db)
			if got != tC.want {
				t.Errorf("Error AddEvent. %s, got=%v, expected=%v", tC.desc, got, tC.want)
			}

		})
	}
}

func TestDelEvent(t *testing.T) {

	var eventTestDel model.Event
	var delID string

	for _, val := range testStore.db {
		eventTestDel = val
		delID = eventTestDel.ID.String()
		break
	}
	countEvent := len(testStore.db)

	testCases := []struct {
		desc string
		id   string
		err  error
		want int
	}{
		{
			desc: "Тест удаления события",
			id:   delID,
			err:  nil,
			want: countEvent - 1,
		},
		{
			desc: "Тест удаления события (должна быть ошибка)", //нет такого события
			id:   delID,
			err:  errors.ErrDelEvent,
			want: countEvent - 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			if _, err := testStore.DelEvent(delID); err != nil {
				if !errors.Is(err, tC.err) {
					t.Errorf("%s error: %v", tC.desc, err)
				}
			}

			got := len(testStore.db)
			if got != tC.want {
				t.Errorf("%s, got=%v, expected=%v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestUpdateEvent(t *testing.T) {

	var eventTest model.Event
	for _, val := range testStore.db {
		eventTest = val
		break
	}

	testCases := []struct {
		desc   string
		event  model.Event
		toText string
		want   string
		err    error
	}{
		{
			desc:   "Тест 1(изменяем Text)",
			event:  eventTest,
			toText: "Meeting",
			want:   "Meeting",
			err:    nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			tC.event.Text = tC.toText
			if _, err := testStore.UpdateEvent(&tC.event); err != nil {
				if !errors.Is(err, tC.err) {
					t.Errorf("%s error: %v", tC.desc, err)
				}
			}
			got, err := testStore.FindEventByID(eventTest.ID.String())

			if got.Text != tC.want || err != nil {
				t.Errorf("%s, got=%v, expected=%v", tC.desc, got.Text, tC.want)
			}
		})
	}
}
