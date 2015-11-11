package gosimplelogger

import (
	"fmt"
	"log"
	"os"
)

const (
	LogPanic = iota
	LogFatal
	LogError
	LogInfo
	LogVerbose
	LogDebug
)

var (
	LogLevel      = LogInfo
	verboseLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func init() {
	verboseLogger = log.New(os.Stdout, "", 0)
	infoLogger = log.New(os.Stdout, "", 0)
	errorLogger = log.New(os.Stderr, "", 0)
}

func logf(logger *log.Logger, logLevel int, format string, v ...interface{}) {
	if LogLevel >= logLevel {
		logger.Println(fmt.Sprintf(format, v...))
	}
}

func logln(logger *log.Logger, logLevel int, v ...interface{}) {
	logf(logger, logLevel, fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	logln(verboseLogger, LogDebug, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	logf(verboseLogger, LogDebug, format, v...)
}

func Verbose(v ...interface{}) {
	logln(verboseLogger, LogVerbose, fmt.Sprint(v...))
}

func Verbosef(format string, v ...interface{}) {
	logf(verboseLogger, LogVerbose, format, v...)
}

func Info(v ...interface{}) {
	logln(infoLogger, LogInfo, v...)
}

func Infof(format string, v ...interface{}) {
	logf(infoLogger, LogInfo, format, v...)
}

func Error(v ...interface{}) {
	logln(errorLogger, LogError, v...)
}

func Errorf(format string, v ...interface{}) {
	logf(errorLogger, LogError, format, v...)
}

func Fatal(v ...interface{}) {
	logln(errorLogger, LogFatal, v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	logf(errorLogger, LogFatal, format, v...)
	os.Exit(1)
}

func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	logln(errorLogger, LogPanic, s)
	panic(s)
}

func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	logln(errorLogger, LogPanic, s)
	panic(s)
}
