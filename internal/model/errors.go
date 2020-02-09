package model

import (
	"errors"
)

var (
	//ErrNotEvent ошибка "событие не найдено"
	ErrNotEvent = errors.New("Событие не найдено")
	//ErrAddEvent "ошибка добавления события"
	ErrAddEvent = errors.New("Ошибка добавления события")
	//ErrDelEvent "ошибка удаления события"
	ErrDelEvent = errors.New("Ошибка удаления события")
	//ErrEditEvent "ошибка изменения события"
	ErrEditEvent = errors.New("Ошибка изменения события")
)
