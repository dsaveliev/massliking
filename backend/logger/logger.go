package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger

	LogFile *os.File
)

func Init(logPath string) {
	LogFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("Error open log file: %v", err)
	}

	Info = log.New(
		LogFile,
		"[ INFO] ",
		log.Ldate|log.Ltime|log.Lmicroseconds,
	)
	Warn = log.New(
		LogFile,
		"[ WARN] ",
		log.Ldate|log.Ltime|log.Lmicroseconds,
	)
	Error = log.New(
		LogFile,
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lmicroseconds,
	)
}

func TaggedLoggers(packageName string, functionName string, args ...string) (
	func(data string), func(data string), func(data string, err error)) {

	tags := fmt.Sprintf("[%s][%s]", packageName, functionName)
	for _, a := range args {
		tags = fmt.Sprintf("%s[%s]", tags, a)
	}

	logInfo := func(data string) {
		Info.Printf("%s %s\n", tags, data)
	}
	logWarn := func(data string) {
		Warn.Printf("%s %s\n", tags, data)
	}
	logError := func(data string, err error) {
		Error.Printf("%s %s: %s\n", tags, data, err)
	}

	return logInfo, logWarn, logError
}

func Stop() {
	err := LogFile.Close()

	if err != nil {
		log.Fatalf("Error closing file: %v", err)
	}
}

func Logger(c *gin.Context) {
	start := time.Now()

	c.Next()

	Info.Printf(
		"| %6s %d | %12s | %s",
		c.Request.Method,
		c.Writer.Status(),
		time.Since(start),
		c.Request.URL.Path,
	)
}
