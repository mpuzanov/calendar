package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/calendar"
	"github.com/mpuzanov/otus-go/calendar/internal/model"
)

//sampleWorkCalendar пример работы с календарём
func sampleWorkCalendar(cal *calendar.Calendar) {
	var err error

	fmt.Println("================================================")
	fmt.Println("Проверяем работу календаря")

	user := "user1"
	startTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 09:00")
	endTime, _ := time.Parse("2006-01-02 15:04", "2020-04-01 10:30")
	fmt.Println("Добавляем <Событие 1>")
	eventID, err := cal.AddEvent(&model.Event{Header: "Событие 1", StartTime: startTime, EndTime: endTime, UserName: user})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Добавляем <Событие 2>")
	_, err = cal.AddEvent(&model.Event{Header: "Событие 2", StartTime: startTime, EndTime: endTime, UserName: user})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(cal.GetUserEvents(user))

	fmt.Println("Удаляем событие c UUID:", eventID)
	_, err = cal.DelEvent(eventID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Показываем результат:")
	fmt.Println(cal.GetUserEvents("user1"))
	fmt.Println("================================================")
}
