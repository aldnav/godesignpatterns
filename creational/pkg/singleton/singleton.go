package creational

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	loglevel int
}

func (l *MyLogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

func (l *MyLogger) SetLogLevel(level int) {
	l.loglevel = level
}

var logger *MyLogger
var once sync.Once

func GetLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("Creating new logger")
		logger = &MyLogger{}
	})
	fmt.Println("Returning existing logger")
	return logger
}
