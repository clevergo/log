// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package zapadapter

import (
	"go.uber.org/zap"
)

// Logger is a zap logger adapter that implements logger interface.
type Logger struct {
	*zap.SugaredLogger
}

// New returns a logger.
func New(sugaredLogger *zap.SugaredLogger) *Logger {
	return &Logger{
		SugaredLogger: sugaredLogger,
	}
}

// Debug implements Logger.Debug.
func (l *Logger) Debug(args ...interface{}) {
	l.SugaredLogger.Debug(args...)
}

// Debugf implements Logger.Debugf.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.SugaredLogger.Debugf(format, args...)
}

// Debugln implements Logger.Debugln.
func (l *Logger) Debugln(args ...interface{}) {
	l.SugaredLogger.Debug(args...)
}

// Info implements Logger.Info.
func (l *Logger) Info(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

// Infof implements Logger.Infof.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.SugaredLogger.Infof(format, args...)
}

// Infoln implements Logger.Infoln.
func (l *Logger) Infoln(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

// Warn implements Logger.Warn.
func (l *Logger) Warn(args ...interface{}) {
	l.SugaredLogger.Warn(args...)
}

// Warnf implements Logger.Warnf.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.SugaredLogger.Warnf(format, args...)
}

// Warnln implements Logger.Warnln.
func (l *Logger) Warnln(args ...interface{}) {
	l.SugaredLogger.Warn(args...)
}

// Error implements Logger.Error.
func (l *Logger) Error(args ...interface{}) {
	l.SugaredLogger.Error(args...)
}

// Errorf implements Logger.Errorf.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.SugaredLogger.Errorf(format, args...)
}

// Errorln implements Logger.Errorln.
func (l *Logger) Errorln(args ...interface{}) {
	l.SugaredLogger.Error(args...)
}

// Fatal implements Logger.Fatal.
func (l *Logger) Fatal(args ...interface{}) {
	l.SugaredLogger.Fatal(args...)
}

// Fatalf implements Logger.Fatalf.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.SugaredLogger.Fatalf(format, args...)
}

// Fatalln implements Logger.Fatalln.
func (l *Logger) Fatalln(args ...interface{}) {
	l.SugaredLogger.Fatal(args...)
}
