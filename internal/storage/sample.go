package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/model"
	m "github.com/mpuzanov/otus-go/calendar/internal/storage/memory"
)

//ISample пример работы с календарём через Интерфейс
func ISample() {
	var err error

	fmt.Println("=====================================")
	fmt.Println("Работаем с календарём через Интерфейс")
	fmt.Println("создаём структуру хранения событий")
	err = NewStorage(Memory)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Добавляем событие")
	event := model.Event{Header: "Событие 1", Date: time.Now()}
	err = DB.AddEvent(&event)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AddEvent(m.NewEvent("Событие 2", time.Now()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(DB.GetEvents())

	fmt.Println("Удаляем событие c UUID:", event.UUID)
	err = DB.DelEvent(&event)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(DB.GetEvents())
}

//Sample пример работы с календарём
func Sample() {
	var err error

	fmt.Println("создаём структуру хранения событий")
	listEvent := m.NewEventStore()

	fmt.Println("Добавляем событие")
	event := m.NewEvent("Событие 1", time.Now())
	err = listEvent.AddEvent(event)
	if err != nil {
		log.Fatal(err)
	}
	err = listEvent.AddEvent(m.NewEvent("Событие 2", time.Now()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(listEvent)

	fmt.Println("Удаляем событие c UUID:", event.UUID)
	err = listEvent.DelEvent(event)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(listEvent)
}
