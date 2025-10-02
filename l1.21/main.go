package main

import (
	"fmt"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Debug(message string)
}

type ThirdPartyLogger struct {
	prefix string
}

func (t *ThirdPartyLogger) LogMessage(level int, msg string) {
	levels := []string{"DEBUG", "INFO", "WARN", "ERROR"}
	fmt.Printf("[%s] %s: %s\n", t.prefix, levels[level], msg)
}

type LoggerAdapter struct {
	thirdPartyLogger *ThirdPartyLogger
}

func NewLoggerAdapter(prefix string) *LoggerAdapter {
	return &LoggerAdapter{
		thirdPartyLogger: &ThirdPartyLogger{prefix: prefix},
	}
}

func (l *LoggerAdapter) Info(message string) {
	l.thirdPartyLogger.LogMessage(1, message)
}

func (l *LoggerAdapter) Error(message string) {
	l.thirdPartyLogger.LogMessage(3, message)
}

func (l *LoggerAdapter) Debug(message string) {
	l.thirdPartyLogger.LogMessage(0, message)
}

func main() {
	logger := NewLoggerAdapter("MyApp")
	
	logger.Info("App started")
	logger.Debug("Loading config...")
	logger.Error("Failed to open connection with database")
}

/*
	Когда стоит использовать паттерн:
	- Интеграция сторонних пакетов
	- Работа с легаси-кодом
	- Тестирование (мокирование)

	Плюсы:
	+ Single responsibility principle
	+ Open-closed principle
	+ Code reuse
	+ Упрощает интеграцию с внешними библиотеками
	+ Позволяет работать с легаси-кодом без его изменения

	Минусы:
	- Увеличивает общую сложность кода
	- Иногда проще изменить исходный код, чем создавать адаптер
	- Может ухудшить производительность из-за дополнительного уровня абстракции

	Еще примеры, кроме кода выше:
   	+ Работа с разными форматами данных (JSON, XML, YAML, CSV)
   	+ Работа с разными БД через единый интерфейс
   	+ Работа с разными message brokers (Kafka, RabbitMQ и пр.)
*/