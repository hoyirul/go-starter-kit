package logger

import (
	"log"
	"os"
)

var (
    Info  *log.Logger
    Error *log.Logger
)

func Init() {
    Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
    Info.Println(message)
}

func LogError(message string, err error) {
    Error.Printf("%s: %v\n", message, err)
}
