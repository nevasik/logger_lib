package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func getZeroLogger() (*zerolog.Logger, *zerolog.ConsoleWriter) {
	onceL.Do(func() {
		moscowLocation, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			log.Println("Failed to load Moscow timezone")
		}

		zerolog.ErrorStackMarshaler = getMarshalStack(3)
		zerolog.TimeFieldFormat = time.StampMilli
		zerolog.ErrorFieldName = zerolog.MessageFieldName
		zerolog.TimestampFunc = func() time.Time {
			return time.Now().In(moscowLocation)
		}

		consoleWriterPtr = &zerolog.ConsoleWriter{Out: os.Stderr,
			TimeFormat:    time.TimeOnly,
			FieldsExclude: []string{"stack"},
			NoColor:       false,
		}

		zLogger := zerolog.New(consoleWriterPtr).
			With().
			Timestamp().
			Caller().
			Stack().
			Logger()

		zl = &zLogger
	})

	return zl, consoleWriterPtr
}

func getMarshalStack(frameAmount int) func(err error) interface{} {
	return func(err error) interface{} {
		type stackTracer interface {
			StackTrace() errors.StackTrace
		}
		sterr, ok := err.(stackTracer)
		if !ok {
			return nil
		}

		st := sterr.StackTrace()
		out := make([]map[string]string, 0, 3)

		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		flag := true
		for i := 1; i < min(frameAmount+1, len(st)-1); i++ {
			pwd, line, funcName := stackInfo(st[i])
			localPath, _ := filepath.Rel(wd, pwd)
			if strings.Contains(localPath, "src/runtime") {
				break
			}

			if flag && strings.Contains(localPath, "loggerM.go") {
				flag = false
				continue
			}

			source := fmt.Sprintf("%s:%s", localPath, line)

			out = append(out, map[string]string{
				"source": source,
				"func":   funcName,
			})
		}
		return out
	}
}

func stackInfo(f errors.Frame) (string, string, string) {
	fn := runtime.FuncForPC(pc(f))
	if fn == nil {
		return "unknown", "unknown", "unknown"
	}
	pwd, line := fn.FileLine(pc(f))
	funcName := funcname(fn.Name())

	return pwd, strconv.Itoa(line), funcName
}

func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

func pc(f errors.Frame) uintptr { return uintptr(f) - 1 }
