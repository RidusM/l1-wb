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