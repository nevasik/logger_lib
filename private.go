package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"os"
)

// ----------INFO----------
func (lc *lg) realInfof(msg string, args ...any) {
	lc.wgLogs.Add(1)
	lc.zl.Info().CallerSkipFrame(2).Msgf(msg, args...)
	lc.wgLogs.Done()
}

func (lc *lg) Infof(format string, args ...interface{}) {
	lc.realInfof(format, args)
}

func (lc *lg) Info(msg string) {
	lc.realInfof(msg)
}

// ----------ERROR----------
func (lc *lg) realErrorf(msg string, args ...interface{}) {
	lc.wgLogs.Add(1)
	lc.zl.Error().CallerSkipFrame(2).Err(errors.Errorf(msg, args...)).Send()
	lc.wgLogs.Done()
}

func (lc *lg) Errorf(format string, args ...interface{}) {
	lc.realErrorf(format, args...)
}

func (lc *lg) Error(msg string) {
	lc.realErrorf(msg)
}

// ----------TRACE----------
func (lc *lg) realTracef(msg string, args ...any) {
	lc.wgLogs.Add(1)
	lc.zl.Trace().CallerSkipFrame(2).Msgf(msg, args...)
	lc.wgLogs.Done()
}

func (lc *lg) Tracef(msg string, args ...any) {
	lc.realTracef(msg, args...)
}

func (lc *lg) Trace(msg string) {
	lc.realTracef(msg)
}

// ----------PANIC----------
func (lc *lg) realPanicf(msg string, args ...interface{}) {
	lc.wgLogs.Add(1)
	lc.zl.WithLevel(zerolog.PanicLevel).CallerSkipFrame(2).Err(errors.Errorf(msg, args...)).Send()

	lc.wgLogs.Wait()
	close(lc.stopCh)
	lc.wgFlush.Wait()
	panic(msg)
}

func (lc *lg) Panicf(format string, args ...any) {
	lc.realPanicf(format, args...)
}

func (lc *lg) Panic(msg string) {
	lc.realPanicf(msg)
}

// ----------DEBUG----------
func (lc *lg) realDebugf(msg string, args ...any) {
	lc.wgLogs.Add(1)
	lc.zl.Debug().CallerSkipFrame(2).Msgf(msg, args...)
	lc.wgLogs.Done()
}

func (lc *lg) Debugf(format string, args ...interface{}) {
	lc.realDebugf(format, args...)
}

func (lc *lg) Debug(msg string) {
	lc.realDebugf(msg)
}

// ----------FATAL----------
func (lc *lg) realFatalf(msg string, args ...interface{}) {
	lc.wgLogs.Add(1)
	lc.zl.WithLevel(zerolog.FatalLevel).CallerSkipFrame(2).Err(errors.Errorf(msg, args...)).Send()

	lc.wgLogs.Wait()
	close(lc.stopCh)
	lc.wgFlush.Wait()
	os.Exit(1)
}

func (lc *lg) Fatalf(format string, args ...interface{}) {
	lc.realFatalf(format, args...)
}

func (lc *lg) Fatal(msg string) {
	lc.realFatalf(msg)
}

// ----------WARN----------
func (lc *lg) realWarnf(msg string, args ...interface{}) {
	lc.wgLogs.Add(1)
	lc.zl.Warn().CallerSkipFrame(2).Err(errors.Errorf(msg, args...)).Send()
	lc.wgLogs.Done()
}

func (lc *lg) Warnf(format string, args ...interface{}) {
	lc.realWarnf(format, args...)
}

func (lc *lg) Warn(msg string) {
	lc.realWarnf(msg)
}
