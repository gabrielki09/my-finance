package logger

import (
	"log"
	"os"
)

type LoggerGroup struct {
	General *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Success *log.Logger
}

var General *LoggerGroup

func Init() {
	General = newLoggerGroup()
}

func newLoggerGroup() *LoggerGroup {
	return &LoggerGroup{
		General: newLogger("General Logger:\t"),
		Info:    newLogger("Info Logger:\t"),
		Error:   newLogger("Error Logger:\t"),
		Success: newLogger("Success Logger:\t"),
	}
}

func newLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}
