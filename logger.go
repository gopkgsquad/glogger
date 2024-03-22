package glogger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// LogLevel represents the log level.
type LogLevel int

const (
	// LogLevelDebug represents debug log level.
	LogLevelDebug LogLevel = iota
	// LogLevelInfo represents info log level.
	LogLevelInfo
	// LogLevelWarning represents warning log level.
	LogLevelWarning
	// LogLevelError represents error log level.
	LogLevelError
	// LogLevelFatal represents fatal log level.
	LogLevelFatal
)

// ANSI color escape codes
const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorReset   = "\033[0m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
)

// Logger represents a custom logger with different log levels.
type Logger struct {
	level            LogLevel
	loggers          map[LogLevel]*log.Logger
	timeColor        string
	fileColor        string
	renderCallerInfo bool
}

// NewLogger creates a new instance of the custom logger.
func NewLogger(out io.Writer, level LogLevel, renderCallerInfo bool) *Logger {
	loggers := make(map[LogLevel]*log.Logger)
	for _, lvl := range []LogLevel{LogLevelDebug, LogLevelInfo, LogLevelWarning, LogLevelError, LogLevelFatal} {
		loggers[lvl] = log.New(out, "", 0)
	}
	return &Logger{
		level:            level,
		loggers:          loggers,
		timeColor:        ColorCyan,
		fileColor:        ColorMagenta,
		renderCallerInfo: renderCallerInfo,
	}
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string) {
	l.logWithColor(LogLevelDebug, msg)
}

// Info logs an info message.
func (l *Logger) Info(msg string) {
	l.logWithColor(LogLevelInfo, msg)
}

// Warning logs a warning message.
func (l *Logger) Warning(msg string) {
	l.logWithColor(LogLevelWarning, msg)
}

// Error logs an error message.
func (l *Logger) Error(msg string) {
	l.logWithColor(LogLevelError, msg)
}

// Fatal logs a fatal message and exits the application.
func (l *Logger) Fatal(msg string) {
	l.logWithColor(LogLevelFatal, msg)
	os.Exit(1)
}

func (l *Logger) logWithColor(level LogLevel, msg string) {
	if l.level <= level {
		timestamp := fmt.Sprintf("%s[%s]%s ", l.timeColor, time.Now().Format("2006/01/02 15:04:05"), ColorReset)
		levelColor := l.getLevelColor(level)
		logMessage := timestamp + levelColor + msg + ColorReset
		if l.renderCallerInfo {
			caller := l.getCallerInfo()
			callerInfo := fmt.Sprintf("%s[%s]%s ", l.fileColor, caller, ColorReset)
			logMessage = callerInfo + logMessage
		}
		l.loggers[level].Output(3, logMessage) // Adjusted depth to log the correct caller
	}
}

func (l *Logger) getCallerInfo() string {
	_, file, line, _ := runtime.Caller(3) // Adjusted depth to get caller outside logger functions
	file = filepath.Base(file)
	return fmt.Sprintf("%s[%s:%d]%s", l.fileColor, file, line, ColorReset)
}

func (l *Logger) getLevelColor(level LogLevel) string {
	switch level {
	case LogLevelDebug:
		return ColorBlue
	case LogLevelInfo:
		return ColorGreen
	case LogLevelWarning:
		return ColorYellow
	case LogLevelError, LogLevelFatal:
		return ColorRed
	default:
		return ColorReset
	}
}

// Infof logs an info message with formatting.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logWithColorf(LogLevelInfo, format, args...)
}

// Warningf logs a warning message with formatting.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.logWithColorf(LogLevelWarning, format, args...)
}

// Errorf logs an error message with formatting.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logWithColorf(LogLevelError, format, args...)
}

// Fatalf logs a fatal message with formatting and exits the application.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logWithColorf(LogLevelFatal, format, args...)
	os.Exit(1)
}

func (l *Logger) logWithColorf(level LogLevel, format string, args ...interface{}) {
	if l.level <= level {
		message := fmt.Sprintf(format, args...)
		l.logWithColor(level, message)
	}
}
