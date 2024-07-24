package logging

import (
	"fmt"
	"sync"
)

type Logger struct {
}

var (
	loggerInstance *Logger
	loggerOnce     sync.Once
)

func InitializeLogger() {
	loggerOnce.Do(func() {
		loggerInstance = &Logger{}
	})
}

func GetLogger() *Logger {
	return loggerInstance
}

func (l *Logger) Info(msg string) {
	fmt.Printf("INFO: %s\n", msg)
}

func (l *Logger) Debug(msg string) {
	fmt.Printf("DEBUG: %s\n", msg)
}

func (l *Logger) Warning(msg string) {
	fmt.Printf("Warning: %s\n", msg)
}

func (l *Logger) Error(msg string) {
	fmt.Printf("ERROR: %s\n", msg)
}
