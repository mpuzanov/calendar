package memslice

import (
	"strconv"
	"testing"
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/errors"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

var (
	testStore          *EventStore
	eventTest          model.Event
	startTime, endTime time.Time
	userName           string
)

func init() {
	generateTestData()
	//fmt.Println(testStore.GetEvents())
}

func generateTestData() {
	userName = "user1"
	startTime = time.Date(2020, time.April, 1, 9, 0, 0, 0, time.UTC)
	endTime, _ = time.Parse("2006-01-02 15:04", "2020-04-01 10:30")
	testStore = NewEventStore()
	for i := 1; i <= 10; i++ {
		testStore.CreateEvent(userName, "Event "+strconv.Itoa(i), "", startTime, endTime)
	}
}

func TestFindEventByID(t *testing.T) {

	testCases := []struct {
		desc string
		find string
		want string
		err  error
	}{
		{
			desc: "Тест 1",
			find: testStore.db[1].ID.String(),
			want: testStore.db[1].ID.String(),
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
	eventTest = model.Event{User: userName, Header: "Event Add", StartTime: startTime, EndTime: endTime}
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
		{
			desc:        "Test (error)", // Header совпадает
			eventHeader: "Event Add",
			err:         errors.ErrAddEvent,
			want:        countEvent + 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			if err := testStore.AddEvent(&eventTest); err != nil {
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

	countEvent := len(testStore.db)
	//Добавляем событие для удаления
	eventTestDel = *NewEvent("user1", "Event del_1", "", startTime, endTime)
	testStore.db = append(testStore.db, eventTestDel)

	testCases := []struct {
		desc     string
		eventDel model.Event
		err      error
		want     int
	}{
		{
			desc:     "Тест удаления события",
			eventDel: eventTestDel,
			err:      nil,
			want:     countEvent,
		},
		{
			desc:     "Тест удаления события (должна быть ошибка)", //нет такого события
			eventDel: eventTestDel,
			err:      errors.ErrDelEvent,
			want:     countEvent,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			if err := testStore.DelEvent(&tC.eventDel); err != nil {
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

	eventTest = testStore.db[0]

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
			if err := testStore.UpdateEvent(&tC.event); err != nil {
				if !errors.Is(err, tC.err) {
					t.Errorf("%s error: %v", tC.desc, err)
				}
			}
			got := testStore.db[0].Text
			if got != tC.want {
				t.Errorf("%s, got=%v, expected=%v", tC.desc, got, tC.want)
			}
		})
	}
}
