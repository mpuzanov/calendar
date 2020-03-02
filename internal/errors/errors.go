package errors

import "errors"

//EventError для выдачи ошибок по событиям календаря
type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	//ErrNotEvent ошибка "событие не найдено"
	ErrNotEvent = EventError("Событие не найдено")
	//ErrAddEvent "ошибка добавления события"
	ErrAddEvent = EventError("Ошибка добавления события")
	//ErrDelEvent "ошибка удаления события"
	ErrDelEvent = EventError("Ошибка удаления события")
	//ErrEditEvent "ошибка изменения события"
	ErrEditEvent = EventError("Ошибка изменения события")
	//ErrNoDBAffected ошибка "Действие не затронуло ни одной строки"
	ErrNoDBAffected = EventError("Действие не затронуло ни одной строки")
	//ErrRecordNotFound ошибка "Запись не найдена"
	ErrRecordNotFound = errors.New("Запись не найдена")
)

//Is обёртка над errors.Is
func Is(err, target error) bool {
	return errors.Is(err, target)
}
