package logger

import (
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/config"
	"log"
	"os"
	"time"
)

var StartTime time.Time

func PrintLog(message string) {

	logFileDir := config.Load().LogFileDir

	logFile, err := os.OpenFile(logFileDir+"logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening logFile: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println(message, time.Since(StartTime))
}
