package lg

import (
	"github.com/rs/zerolog"
	"sync"
)

type Logger interface {
	Error(msg string)
	Errorf(msg string, args ...any)
	Trace(msg string)
	Tracef(msg string, args ...any)
	Debug(msg string)
	Debugf(msg string, args ...any)
	Info(msg string)
	Infof(msg string, args ...any)
	Warn(msg string)
	Warnf(msg string, args ...any)
	Fatal(msg string)
	Fatalf(msg string, args ...any)
	Panic(msg string)
	Panicf(msg string, args ...any)
}

var (
	zl               *zerolog.Logger
	consoleWriterPtr *zerolog.ConsoleWriter
	onceL            sync.Once
	logger           *lg
	once             sync.Once
)

type lg struct {
	zl   *zerolog.Logger
	logs *struct {
		logs chan []byte
	}
	wgLogs  sync.WaitGroup
	wgFlush sync.WaitGroup
	mu      sync.Mutex
	stopCh  chan struct{}
}
