// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package log

import (
	"log"
)

// Logger is a general logger interface.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

var _ Logger = New()

// StdLogger is a wrapper of standard logger.
type StdLogger struct {
}

// New returns a standard logger.
func New() *StdLogger {
	return &StdLogger{}
}

// Debug implements Logger's Debug.
func (l *StdLogger) Debug(args ...interface{}) {
	log.Print(args...)
}

// Debugf implements Logger's Debugf.
func (l *StdLogger) Debugf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Info implements Logger's Info.
func (l *StdLogger) Info(args ...interface{}) {
	log.Print(args...)
}

// Infof implements Logger's Infof.
func (l *StdLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Warn implements Logger's Warn.
func (l *StdLogger) Warn(args ...interface{}) {
	log.Print(args...)
}

// Warnf implements Logger's Warnf.
func (l *StdLogger) Warnf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Error implements Logger's Error.
func (l *StdLogger) Error(args ...interface{}) {
	log.Print(args...)
}

// Errorf implements Logger's Errorf.
func (l *StdLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Fatal implements Logger's Fatal.
func (l *StdLogger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf implements Logger's Fatalf.
func (l *StdLogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
