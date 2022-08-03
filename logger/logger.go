package logger

import (
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
		fatalLogger:   log.New(w, "[FATAL]", log.Ldate|log.Ltime|log.Lshortfile),
		silentLogger:  log.New(w, "", 0),
		printlnLogger: log.New(w, "", 0),
		errorLogger:   log.New(w, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:    log.New(w, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(w, "[WARNING]", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger:   log.New(w, "[DEBUG]", log.Ldate|log.Ltime|log.Lshortfile),
		verboseLogger: log.New(w, "[VERBOSE]", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

var stdLogger = NewLogger(LVerbose)

// func (logger *Logger) SetLevel(level int32) {
// 	if level < LFATAL || level > LVerbose {
// 		return
// 	}

// 	atomic.StoreInt32(&logger.level, level)
// }

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

// func SetLevel(level int32) {
// 	if level < LFATAL || level > LVerbose {
// 		return
// 	}

// 	atomic.StoreInt32(&stdLogger.level, level)
// }

// func (logger *Logger) Println(arg ...interface{}) {
// 	if atomic.LoadInt64(&logger.level) < LPrintln {
// 		return
// 	}
// 	logger.printlnLogger.Println(arg...)
// }

// func (logger *Logger) Error(err error) {
// 	if atomic.LoadInt64(&logger.level) < LERROR {
// 		return
// 	}
// 	logger.printlnLogger.Println(err)
// }

// func (logger *Logger) Silent(arg ...interface{}) {
// 	if atomic.LoadInt64(&logger.level) < LSilent {
// 		return
// 	}
// 	logger.silentLogger.Println(arg...)
// }

// func (logger *Logger) Info(str string, arg ...interface{}) {
// 	if atomic.LoadInt64(&logger.level) < LINFO {
// 		return
// 	}
// 	logger.infoLogger.Printf(str, arg...)
// }

func Info(str string, arg ...interface{}) {
	if atomic.LoadInt32(&stdLogger.level) < LINFO {
		return
	}
	stdLogger.infoLogger.Printf(str, arg...)
}

func Error(err error) {
	if atomic.LoadInt32(&stdLogger.level) < LERROR {
		return
	}
	stdLogger.printlnLogger.Println(err)
}

func Println(arg ...interface{}) {
	if atomic.LoadInt32(&stdLogger.level) < LPrintln {
		return
	}
	stdLogger.printlnLogger.Println(arg...)
}

func Silent(arg ...interface{}) {
	if atomic.LoadInt32(&stdLogger.level) < LSilent {
		return
	}
	stdLogger.silentLogger.Println(arg...)
}

// func Info(str string, arg ...interface{}) {
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	log.Printf(str+"\n", arg...)
// }

// func Println(str string, arg ...interface{}) {
// 	log.SetFlags(0)
// 	log.Printf(str+"\n", arg...)
// }
