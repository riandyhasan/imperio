package logger

import (
	"context"
	"encoding/json"
	"os"
	"time"
)

// logger is a concrete implementation of Logger that outputs JSON to stdout.
type logger struct {
	out *os.File
}

// Field represents a key-value pair for structured logs.
type field struct {
	Key   string
	Value interface{}
}

// NewLogger returns a new JSON-formatted logger instance.
func NewLogger() *logger {
	return &logger{out: os.Stdout}
}

type logEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Error     string                 `json:"error,omitempty"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// Info logs a structured info message.
func (l *logger) Info(ctx context.Context, msg string, fields ...field) {
	l.log("INFO", msg, "", fields)
}

// Error logs a structured error message.
func (l *logger) Error(ctx context.Context, msg string, err error, fields ...field) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	l.log("ERROR", msg, errMsg, fields)
}

// log is the internal method to write the log entry.
func (l *logger) log(level, msg, errMsg string, fields []field) {
	entry := logEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   msg,
		Error:     errMsg,
		Fields:    make(map[string]interface{}),
	}

	for _, f := range fields {
		entry.Fields[f.Key] = f.Value
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		os.Stderr.WriteString("Failed to marshal log entry: " + err.Error() + "\n")
		return
	}

	l.out.Write(append(jsonData, '\n'))
}
