package logging

import (
	"io"
	"log"
	"os"
)

var (
	E *log.Logger
	I *log.Logger
	R *FilteredLogger
)

func init() {
	// Ensure the directory exists (MkdirAll is no-op if directory already exists)
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Create log files or open them if they already exist
	errorFile, err := os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening error log file: %v", err)
	}

	infoFile, err := os.OpenFile("logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening info log file: %v", err)
	}

	requestFile, err := os.OpenFile("logs/request.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening request log file: %v", err)
	}

	// Initialize ErrorLogger with the MultiWriter

	// Create custom loggers
	I = log.New(infoFile, "INFO: ", log.LstdFlags)
	// R = log.New(requestFile, "REQUEST: ", log.LstdFlags)
	//for errors i am writing on console and on the file
	multiWriter := io.MultiWriter(os.Stderr, errorFile)
	E = log.New(multiWriter, "ERROR: ", log.LstdFlags)

	baseRequestLogger := log.New(requestFile, "REQUEST: ", log.LstdFlags)
	R = NewFilteredLogger(baseRequestLogger)
	simpleLog()
}

func simpleLog() {

	// Open or create the log file for appending, create it if it doesn't exist
	logFile, err := os.OpenFile("logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Set the global log output to the file
	multiWriter := io.MultiWriter(os.Stderr, logFile)
	log.SetOutput(multiWriter)

	// Optional: Set the log to also output the date and time
	log.SetFlags(log.LstdFlags)
}
