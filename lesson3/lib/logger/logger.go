package logger

import (
	"log"
	"runtime"
	"strings"

	"github.com/misaliperver/golesson/lesson3/config"
)

var (
	env, _ = config.Get()
	levels = map[string]int{
		"silly":   0,
		"debug":   1,
		"info":    2,
		"warning": 3,
		"error":   4,
	}
	currentLevel int = levels[env.ConsoleLevel]
)

type Logger struct {
	Name string
	Path string
}

func NewLogger() *Logger {
	pc, _, _, _ := runtime.Caller(1)

	funcName := runtime.FuncForPC(pc).Name()

	lastSlash := strings.LastIndexByte(funcName, '/')

	if lastSlash < 0 {
		lastSlash = 0
	} else {
		lastSlash++
	}

	// lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash

	return &Logger{
		Name: funcName[lastSlash:] + "()",
		Path: funcName,
	}
}

func checkUsableLogger(level int) bool {
	if currentLevel >= level {
		return false
	}

	return true
}

func (l *Logger) Silly(message ...any) {
	if checkUsableLogger(levels["debug"]) {
		log.Println("SILLY", l.Name, "-", message)
	}
}

func (l *Logger) Debug(message ...any) {
	if checkUsableLogger(levels["debug"]) {
		log.Println("DEBUG", l.Name, "-", message)
	}
}

func (l *Logger) Info(message ...any) {
	if checkUsableLogger(levels["info"]) {
		log.Println("INFO", l.Name, "-", message)
	}
}

func (l *Logger) Warning(message ...any) {
	if checkUsableLogger(levels["warning"]) {
		log.Println("WARN", l.Name, "-", message, "\n          PWD:", l.Path)
	}
}

func (l *Logger) Error(message ...any) {
	if checkUsableLogger(levels["error"]) {
		log.Println("ERROR", l.Name, "-", message, "\n          PWD:", l.Path)
	}
}
