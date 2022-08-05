package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

const (
	LFATAL = 1 << iota
	LSilent
	LPrintln
	LERROR
	LINFO
	LWARNING
	LDEBUG
	LVerbose
)

type Logger struct {
	w             io.Writer
	level         int32
	fatalLogger   *log.Logger
	silentLogger  *log.Logger
	printlnLogger *log.Logger
	errorLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	debugLogger   *log.Logger
	verboseLogger *log.Logger
	mu            sync.Mutex
}

func NewLogger(level int32) *Logger {
	w := os.Stderr
	return &Logger{
		w:             w,
		level:         level,
		fatalLogger:   log.New(w, "[FATAL]", log.Ldate|log.Ltime|log.Llongfile),
		silentLogger:  log.New(w, "", 0),
		printlnLogger: log.New(w, "", 0),
		errorLogger:   log.New(w, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile),
		infoLogger:    log.New(w, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(w, "[WARNING]", log.Ldate|log.Ltime|log.Llongfile),
		debugLogger:   log.New(w, "[DEBUG]", log.Ldate|log.Ltime|log.Llongfile),
		verboseLogger: log.New(w, "[VERBOSE]", log.Ldate|log.Ltime|log.Llongfile),
	}
}

var stdLogger = NewLogger(LVerbose)

func (logger *Logger) SetLevel(level int32) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.level = level
}

func SetLevel(level int32) {
	if level < LFATAL || level > LVerbose {
		return
	}

	stdLogger.SetLevel(level)
}

func Info(arg ...interface{}) {
	if atomic.LoadInt32(&stdLogger.level) < LINFO {
		return
	}
	stdLogger.infoLogger.Output(2, fmt.Sprintln(arg...))
}

func Error(err error) {
	if atomic.LoadInt32(&stdLogger.level) < LERROR {
		return
	}
	stdLogger.errorLogger.Output(2, fmt.Sprintln(err))
}

func Println(arg ...interface{}) {
	if atomic.LoadInt32(&stdLogger.level) < LPrintln {
		return
	}
	stdLogger.printlnLogger.Output(2, fmt.Sprintln(arg...))
}

func Silent(arg ...interface{}) {
	if atomic.LoadInt32(&stdLogger.level) < LSilent {
		return
	}
	stdLogger.silentLogger.Output(2, fmt.Sprintln(arg...))
}
