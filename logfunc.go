package logrus

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

const (
	logCalldepth = 2 // specific to this wrapper implementation
	callerInfo   = "callerInfo"
)

// Alias types from logrus to this package for easier access
type Fields logrus.Fields

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *logrus.Entry {
	f := fileInfo(logCalldepth)
	return log.WithFields(logrus.Fields{
		logrus.ErrorKey: err,
		callerInfo:      f,
	})
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	f := fileInfo(logCalldepth)
	return log.WithFields(logrus.Fields{
		callerInfo: f,
		key:        value,
	})
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *logrus.Entry {
	f := fileInfo(logCalldepth)
	data := logrus.Fields{callerInfo: f}
	for k, v := range fields {
		data[k] = v
	}
	return &logrus.Entry{Logger: log, Data: data}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	f := fileInfo(logCalldepth)
	log.WithField(callerInfo, f).Fatalf(format, args...)
}

// fileInfo grabs the file name and line number of the caller.
func fileInfo(callDepth int) string {
	// Inspect runtime call stack
	pc := make([]uintptr, callDepth)
	runtime.Callers(callDepth, pc)
	f := runtime.FuncForPC(pc[callDepth-1])
	file, line := f.FileLine(pc[callDepth-1])

	// Truncate abs file path
	if slash := strings.LastIndex(file, "/"); slash >= 0 {
		file = file[slash+1:]
	}

	// Truncate package name
	funcName := f.Name()
	if slash := strings.LastIndex(funcName, "."); slash >= 0 {
		funcName = funcName[slash+1:]
	}

	return fmt.Sprintf("%s:%d %s -", file, line, funcName)
}
