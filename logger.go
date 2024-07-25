package logger

import (
	"fmt"
	"os"
	"time"
)


var logLevel = map[int]string{
	0: "Info",
	1: "Warning",
	2: "Error",
	3: "Debug", 
}

type Logger struct {
	level int
	console bool
	file *os.File
}

func NewLogger(level int, console bool, filePath string) (*Logger ,error) {
	logger := &Logger{level: level, console: console}
	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		logger.file = file
	}
	return logger, nil
}

// log writes the log message to the appropriate outputs.
func (l *Logger) log(level int, message string) {
    if level > l.level {
        return
    }
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    logMessage := fmt.Sprintf("%s [%s] %s\n", timestamp, logLevel[level], message)
    
    if l.console {
        fmt.Print(logMessage)
    }
    if l.file != nil {
        l.file.WriteString(logMessage)
    }
}

// Info logs an informational message.
func (l *Logger) Info(message string) {
    l.log(0, message)
}

// Warning logs a warning message.
func (l *Logger) Warning(message string) {
    l.log(1, message)
}

// Error logs an error message.
func (l *Logger) Error(message string) {
    l.log(2, message)
}

// Debug logs a debug message.
func (l *Logger) Debug(message string) {
    l.log(3, message)
}

// Close closes the file if it's open.
func (l *Logger) Close() {
    if l.file != nil {
        l.file.Close()
    }
}