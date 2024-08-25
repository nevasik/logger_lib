package lg

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func Init() Logger {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})

	once.Do(func() {
		zLog, consoleWriter := getZeroLogger()

		newLogger := zLog.Output(consoleWriter)

		logger = &lg{zl: &newLogger, stopCh: make(chan struct{})}
	})

	return logger
}

func Print(msg string) {
	logger.realDebugf(msg)
}

func Printf(format string, args ...any) {
	logger.realDebugf(format, args...)
}

func Println(msg string) {
	logger.realDebugf(msg)
}

func Trace(msg string) {
	logger.realTracef(msg)
}

func Debug(msg string) {
	logger.realDebugf(msg)
}

func Info(msg string) {
	logger.realInfof(msg)
}

func Warn(msg string) {
	logger.realWarnf(msg)
}

func Error(err error) {
	if err != nil {
		logger.realErrorf(err.Error())
	}
}

func Fatal(msg string) {
	logger.realFatalf(msg)
}

func Panic(msg string) {
	logger.realPanicf(msg)
}

func Tracef(format string, args ...any) {
	logger.realTracef(format, args...)
}

func Debugf(format string, args ...any) {
	logger.realDebugf(format, args...)
}

func Infof(format string, args ...any) {
	logger.realInfof(format, args...)
}

func Warnf(format string, args ...any) {
	logger.realWarnf(format, args...)
}

func Errorf(format string, args ...any) {
	logger.realErrorf(format, args...)
}

func Fatalf(format string, args ...any) {
	logger.realFatalf(format, args...)
}

func Panicf(format string, args ...any) {
	logger.realPanicf(format, args...)
}

func (lc *lg) Write(p []byte) (n int, err error) {
	cp := make([]byte, len(p))
	copy(cp, p)

	lc.logs.logs <- cp
	return len(cp), nil
}
