// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package zapadapter

import (
	"github.com/clevergo/log"
	"go.uber.org/zap"
)

// Logger is a zap logger adapter that implements logger interface.
type logger struct {
	*zap.SugaredLogger
}

// New returns a logger.
func New(sugaredLogger *zap.SugaredLogger) log.Logger {
	return &logger{
		SugaredLogger: sugaredLogger,
	}
}

// Debug implements Logger.Debug.
func (l *logger) Debug(args ...interface{}) {
	l.SugaredLogger.Debug(args...)
}

// Debugf implements Logger.Debugf.
func (l *logger) Debugf(format string, args ...interface{}) {
	l.SugaredLogger.Debugf(format, args...)
}

// Debugln implements Logger.Debugln.
func (l *logger) Debugln(args ...interface{}) {
	l.SugaredLogger.Debug(args...)
}

// Fatal implements Logger.Fatal.
func (l *logger) Fatal(args ...interface{}) {
	l.SugaredLogger.Fatal(args...)
}

// Fatalf implements Logger.Fatalf.
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.SugaredLogger.Fatalf(format, args...)
}

// Fatalln implements Logger.Fatalln.
func (l *logger) Fatalln(args ...interface{}) {
	l.SugaredLogger.Fatal(args...)
}

// Info implements Logger.Info.
func (l *logger) Info(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

// Infof implements Logger.Infof.
func (l *logger) Infof(format string, args ...interface{}) {
	l.SugaredLogger.Infof(format, args...)
}

// Infoln implements Logger.Infoln.
func (l *logger) Infoln(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

// Warn implements Logger.Warn.
func (l *logger) Warn(args ...interface{}) {
	l.SugaredLogger.Warn(args...)
}

// Warnf implements Logger.Warnf.
func (l *logger) Warnf(format string, args ...interface{}) {
	l.SugaredLogger.Warnf(format, args...)
}

// Warnln implements Logger.Warnln.
func (l *logger) Warnln(args ...interface{}) {
	l.SugaredLogger.Warn(args...)
}

// Error implements Logger.Error.
func (l *logger) Error(args ...interface{}) {
	l.SugaredLogger.Error(args...)
}

// Errorf implements Logger.Errorf.
func (l *logger) Errorf(format string, args ...interface{}) {
	l.SugaredLogger.Errorf(format, args...)
}

// Errorln implements Logger.Errorln.
func (l *logger) Errorln(args ...interface{}) {
	l.SugaredLogger.Error(args...)
}
