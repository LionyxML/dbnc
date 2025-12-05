package logx

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// --- Internal constants and types ---

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
)

type level int

const (
	levelDebug level = iota
	levelInfo
	levelWarn
	levelError
)

func parseLevel(levelStr string) level {
	switch strings.ToLower(levelStr) {
	case "debug":
		return levelDebug
	case "info":
		return levelInfo
	case "warn", "warning":
		return levelWarn
	case "error":
		return levelError
	default:
		return levelInfo
	}
}

// --- Interface Definition ---

// Logger defines the interface for logging operations.
type Logger interface {
	// Logs in Debug level
	Debug(format string, v ...any)

	// Logs in Info level
	Info(format string, v ...any)

	// Logs in Warn level
	Warn(format string, v ...any)

	// Logs in Error level
	Error(format string, v ...any)
}

// --- Implementation ---

type loggerImpl struct {
	logger   *log.Logger
	level    level
	useColor bool
}

// New creates a new Logger instance.
func New(tag string, levelStr string, useColor bool) Logger {
	logger := log.New(os.Stdout, tag+" ", log.LstdFlags)

	return &loggerImpl{
		logger:   logger,
		level:    parseLevel(levelStr),
		useColor: useColor,
	}
}

func (l *loggerImpl) colorize(color, msg string) string {
	if l.useColor {
		return color + msg + colorReset
	}
	return msg
}

func (l *loggerImpl) Debug(format string, v ...any) {
	if l.level <= levelDebug {
		msg := fmt.Sprintf(format, v...)
		l.logger.Print(l.colorize(colorCyan, "DEBUG: "+msg))
	}
}

func (l *loggerImpl) Info(format string, v ...any) {
	if l.level <= levelInfo {
		msg := fmt.Sprintf(format, v...)
		l.logger.Print(l.colorize(colorBlue, "INFO: "+msg))
	}
}

func (l *loggerImpl) Warn(format string, v ...any) {
	if l.level <= levelWarn {
		msg := fmt.Sprintf(format, v...)
		l.logger.Print(l.colorize(colorYellow, "WARN: "+msg))
	}
}

func (l *loggerImpl) Error(format string, v ...any) {
	if l.level <= levelError {
		msg := fmt.Sprintf(format, v...)
		l.logger.Print(l.colorize(colorRed, "ERROR: "+msg))
	}
}
