package server

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger struct has methods for logging warnings and erros
// while the app is runing
type logger struct {
	file *os.File
}

var (
	//Logger global way to log things
	Logger logger
)

func newLogger() logger {
	f, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("[Error] %v", err)
	}
	log.SetOutput(f)

	return logger{file: f}
}

// Close closes the file of the log instance
func (l *logger) Close() {
	err := l.file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// Add writes the msg to log file
func (l logger) Add(logmsg string) {
	log.Println(fmt.Sprintf("[SERVER %s] %s", time.Now().Format(time.RFC850), logmsg))
}
