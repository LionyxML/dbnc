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
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Warn(format string, v ...any)
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

// Logs in Debug level
func (l *loggerImpl) Debug(format string, v ...any) {
	if l.level <= levelDebug {
		msg := fmt.Sprintf(format, v...)
		l.logger.Print(l.colorize(colorCyan, "DEBUG: "+msg))
	}
}

// Logs in Info level
func (l *loggerImpl) Info(format string, v ...any) {
	if l.level <= levelInfo {
		msg := fmt.Sprintf(format, v...)
		l.logger.Print(l.colorize(colorBlue, "INFO: "+msg))
	}
}

// Logs in Warn level
func (l *loggerImpl) Warn(format string, v ...any) {
	if l.level <= levelWarn {
		msg := fmt.Sprintf(format, v...)
		l.logger.Print(l.colorize(colorYellow, "WARN: "+msg))
	}
}

// Logs in Error level
func (l *loggerImpl) Error(format string, v ...any) {
	if l.level <= levelError {
		msg := fmt.Sprintf(format, v...)
		l.logger.Print(l.colorize(colorRed, "ERROR: "+msg))
	}
}
