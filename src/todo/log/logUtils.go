package log

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("checklist.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// for logging
func Danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	print(false, args...)
}

func Fatal(args ...interface{}) {
	logger.SetPrefix("FATAL ")
	print(true, args...)
}

func Info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	print(false, args...)
}

func Warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	print(false, args...)
}

func print(isFatal bool, args ...interface{}) {
	fmt.Println(args...)

	if isFatal == true {
		logger.Fatalln(args...)
	} else {
		logger.Println(args...)
	}
}
