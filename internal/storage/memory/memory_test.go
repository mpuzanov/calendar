package memory_test

import (
	"errors"
	"testing"
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/model"
	m "github.com/mpuzanov/otus-go/calendar/internal/storage/memory"
)

func TestFindEventByHeader(t *testing.T) {
	testStore := m.EventStore{Events: make([]model.Event, 0)}
	testStore.Events = append(testStore.Events, model.Event{Header: "Event 1", Date: time.Now()})
	testStore.Events = append(testStore.Events, model.Event{Header: "Event 2", Date: time.Now()})

	testCases := []struct {
		desc string
		find string
		want string
		err  error
	}{
		{
			desc: "Тест 1",
			find: "Event 1",
			want: "Event 1",
			err:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := testStore.FindEventByHeader(tC.find)
			if err != nil {
				t.Errorf("%s, got=%v, expected=%v, error: %v", tC.desc, got.Header, tC.want, err)
			}
			if got.Header != tC.want {
				t.Errorf("Error FindEventByHeader. %s, got=%v, expected=%v", tC.desc, got.Header, tC.want)
			}
		})
	}
}
func TestAddEvent(t *testing.T) {
	testStore := m.EventStore{Events: make([]model.Event, 0)}

	testCases := []struct {
		desc string
		to   string
		err  error
		want int
	}{
		{
			desc: "Test (success)",
			to:   "Event 1",
			err:  nil,
			want: 1,
		},
		{
			desc: "Test (error)",
			to:   "Event 1",
			err:  model.ErrAddEvent,
			want: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			event := m.NewEvent(tC.to, time.Now())

			if err := testStore.AddEvent(event); err != nil {
				if !errors.Is(err, tC.err) {
					t.Errorf("%s error: %v", tC.desc, err)
				}
			}
			got := len(testStore.Events)
			if got != tC.want {
				t.Errorf("Error AddEvent. %s, got=%v, expected=%v", tC.desc, got, tC.want)
			}

		})
	}
}

func TestDelEvent(t *testing.T) {
	testStore := m.EventStore{Events: make([]model.Event, 0)}
	event := model.Event{Header: "Event 1", Date: time.Now()}
	testStore.Events = append(testStore.Events, event)

	testCases := []struct {
		desc     string
		eventDel model.Event
		err      error
		want     int
	}{
		{
			desc:     "Тест удаления события",
			eventDel: event,
			err:      nil,
			want:     1,
		},
		{
			desc:     "Тест удаления события (должна быть ошибка)",
			eventDel: model.Event{Header: "Event 1", Date: time.Now()},
			err:      model.ErrDelEvent,
			want:     2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			event := model.Event{Header: "Event X", Date: time.Now()}
			testStore.Events = append(testStore.Events, event)

			if err := testStore.DelEvent(&tC.eventDel); err != nil {
				if !errors.Is(err, tC.err) {
					t.Errorf("%s error: %v", tC.desc, err)
				}
			}

			got := len(testStore.Events)
			if got != tC.want {
				t.Errorf("%s, got=%v, expected=%v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestSetEvent(t *testing.T) {
	testStore := m.EventStore{Events: make([]model.Event, 0)}
	event := model.Event{Header: "Event 1", Date: time.Now(), Description: "first description"}
	testStore.Events = append(testStore.Events, event)

	testCases := []struct {
		desc   string
		event  model.Event
		toDesc string
		want   string
		err    error
	}{
		{
			desc:   "Тест 1(изменяем Description)",
			event:  event,
			toDesc: "Meeting",
			want:   "Meeting",
			err:    nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.event.Description = tC.toDesc
			if err := testStore.SetEvent(&tC.event); err != nil {
				if !errors.Is(err, tC.err) {
					t.Errorf("%s error: %v", tC.desc, err)
				}
			}
			got := testStore.Events[0].Description
			if got != tC.want {
				t.Errorf("%s, got=%v, expected=%v", tC.desc, got, tC.want)
			}
		})
	}
}
