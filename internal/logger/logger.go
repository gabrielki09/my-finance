package logger

import (
	"log"
	"os"
	"path/filepath"
)

type LoggerGroup struct {
	General *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Success *log.Logger
}

var General *LoggerGroup

func Init() {
	General = newLoggerGroup("general")
}

func newLoggerGroup(scope string) *LoggerGroup {
	logFileName := "general.log"

	return &LoggerGroup{
		General: newLogger("General Logger:\t", scope, logFileName),
		Info:    newLogger("Info Logger:\t", scope, logFileName),
		Error:   newLogger("Error Logger:\t", scope, logFileName),
		Success: newLogger("Success Logger:\t", scope, logFileName),
	}
}

func newLogger(prefix, scope, fileName string) *log.Logger {
	appEnv := os.Getenv("APP_ENV")

	if appEnv != "dev" {
		return log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	}

	filePath := filepath.Join("log", scope, fileName)
	logDir := filepath.Dir(filePath)

	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal("Erro ao criar o diretório de log:", err)
	}

	file, err := os.OpenFile(
		filePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		log.Fatal("Erro ao abrir o arquivo de log:", err)
	}

	return log.New(file, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}
