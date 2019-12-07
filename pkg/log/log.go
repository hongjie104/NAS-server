package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Level a
type Level int

var (
	// F a
	F *os.File

	// DefaultPrefix a
	DefaultPrefix = ""
	// DefaultCallerDepth a
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	// Debug a
	Debug Level = iota
	// Info a
	Info
	// Warning a
	Warning
	// Error a
	Error
	// Fatal a
	Fatal
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// LogDebug a
func LogDebug(v ...interface{}) {
	setPrefix(Debug)
	logger.Println(v)
}

// LogInfo a
func LogInfo(v ...interface{}) {
	setPrefix(Info)
	logger.Println(v)
}

func LogWarn(v ...interface{}) {
	setPrefix(Warning)
	logger.Println(v)
}

func LogError(v ...interface{}) {
	setPrefix(Error)
	logger.Println(v)
}

func LogFatal(v ...interface{}) {
	setPrefix(Fatal)
	logger.Println(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
