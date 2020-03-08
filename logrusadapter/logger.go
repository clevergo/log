// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package logrusadapter

import (
	"github.com/sirupsen/logrus"
)

// Logger is a logrus logger adapter that implements logger interface.
type Logger struct {
	*logrus.Logger
}

// New returns a logger.
func New(logrusLogger *logrus.Logger) *Logger {
	return &Logger{logrusLogger}
}
