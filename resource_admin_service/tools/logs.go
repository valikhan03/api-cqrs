package tools

import(
	"os"
	"log"
)

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func InitLogger(path, flag string) *log.Logger {
	logFile, err := OpenLogFile(path)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(logFile, flag, log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	return logger
}
