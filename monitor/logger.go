package monitor

import (
	"fmt"
	"log"
	"os"
	"time"
)

type CustomLogger struct {
	params  string
	logFile *os.File
}

func NewCustomLogger(v ...interface{}) *CustomLogger {
	if len(v) < 1 {
		log.Fatal("NewCustomLogger: At least one argument is required.")
	}

	logFilePath := "log.log"
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}

	var params string

	for _, pd := range v {
		params += fmt.Sprintf("[%s] ", pd)
	}

	return &CustomLogger{
		params:  params,
		logFile: logFile,
	}
}

func getCurrentTimeWithMillis() string {
	return time.Now().Format("15:04:05.000000")
}

func (cl *CustomLogger) Info(v ...interface{}) {
	var message string

	for _, m := range v {
		message += fmt.Sprintf("%s ", m)
	}

	fmt.Printf("[%s] %s%s\n", getCurrentTimeWithMillis(), cl.params, message)
	fmt.Fprintf(cl.logFile, "[%s] %s%s\n", getCurrentTimeWithMillis(), cl.params, message)
}
