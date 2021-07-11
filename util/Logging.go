package util

import (
	"github.com/jgrath/go-and-find-with-go/config"
	"log"
	"os"
	"time"
)

var (
	LogWarning        *log.Logger
	LogInfo           *log.Logger
	LogError          *log.Logger
	loggingOutputFile *os.File
)

const loggingDateTimeFormat = "-01-02-2006-15-04-054"
const infoPrefix = "INFO: "
const errorPrefix = "ERROR: "
const warningPrefix = "WARNING: "

func init() {
	setupLogging()
}

func setupLogging() {

	logConfig := config.GetConfiguration()

	fileNameStamp := time.Now().Format(loggingDateTimeFormat)

	loggingOutputFile, err := os.OpenFile(logConfig.Logfile + fileNameStamp + ".log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	LogInfo = log.New(loggingOutputFile, infoPrefix, log.Ldate|log.Ltime|log.Lshortfile)
	LogWarning = log.New(loggingOutputFile, warningPrefix, log.Ldate|log.Ltime|log.Lshortfile)
	LogError = log.New(loggingOutputFile, errorPrefix, log.Ldate|log.Ltime|log.Lshortfile)
}
