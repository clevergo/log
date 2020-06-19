// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package log

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdLogger(t *testing.T) {
	cases := []struct {
		method string
		format string
		args   []interface{}
		out    string
	}{
		{"debug", "", []interface{}{"debug"}, "debug"},
		{"debugf", "msg: %s", []interface{}{"debugf"}, "msg: debugf"},
		{"info", "", []interface{}{"info"}, "info"},
		{"infof", "msg: %s", []interface{}{"infof"}, "msg: infof"},
		{"warn", "", []interface{}{"warn"}, "warn"},
		{"warnf", "msg: %s", []interface{}{"warnf"}, "msg: warnf"},
		{"error", "", []interface{}{"error"}, "error"},
		{"errorf", "msg: %s", []interface{}{"errorf"}, "msg: errorf"},
	}
	logger := New()
	out := &bytes.Buffer{}
	log.SetOutput(out)
	for _, test := range cases {
		out.Reset()
		switch test.method {
		case "debug":
			logger.Debug(test.args...)
		case "debugf":
			logger.Debugf(test.format, test.args...)
		case "info":
			logger.Info(test.args...)
		case "infof":
			logger.Infof(test.format, test.args...)
		case "warn":
			logger.Warn(test.args...)
		case "warnf":
			logger.Warnf(test.format, test.args...)
		case "error":
			logger.Error(test.args...)
		case "errorf":
			logger.Errorf(test.format, test.args...)
		}
		assert.Contains(t, out.String(), test.out)
	}
}
