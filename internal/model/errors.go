package model

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
)
