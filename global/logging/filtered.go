package logging

import (
	"log"
	"strings"
)

type FilteredLogger struct {
	*log.Logger
}

// NewFilteredLogger creates a new FilteredLogger wrapping the given logger.
func NewFilteredLogger(l *log.Logger) *FilteredLogger {
	return &FilteredLogger{l}
}

// LogIf meets criteria for logging.
func (fl *FilteredLogger) LogIf(message string) {
	// Define your criteria for filtering logs here.
	// Example: Skip messages that are too long or contain "password"
	if len(message) > 1000 || strings.Contains(message, "password") {
		return // Skip logging this message
	}
	// If the message meets your criteria, log it
	fl.Println(message)
}
