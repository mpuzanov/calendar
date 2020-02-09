/*
Package logger Пакет для логирования работы

//Создание логгеров и настройка уровня логирования
configLog := logger.Configuration{Level: "debug", JSONFormat: false}
logr, err := logger.NewLogger(configLog)
if err != nil {
	log.Fatalf("Could not instantiate log %s", err.Error())
}

*/
package logger
