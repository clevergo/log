// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package zapadapter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/clevergo/log"
	"go.uber.org/zap"
)

var (
	tempLogFile1 *os.File
	tempLogFile2 *os.File
	testLogger   log.Logger
	zaplogger    *zap.SugaredLogger
)

func newTestZapLogger(logFile string) (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{logFile}
	cfg.EncoderConfig.TimeKey = ""
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

func TestMain(m *testing.M) {
	var err error
	tempLogFile1, err = ioutil.TempFile("", "test.log")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempLogFile1.Name())
	logger1, err := newTestZapLogger(tempLogFile1.Name())
	if err != nil {
		panic(err)
	}
	testLogger = New(logger1)

	tempLogFile2, err = ioutil.TempFile("", "test.log.2")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempLogFile2.Name())
	zaplogger, err = newTestZapLogger(tempLogFile2.Name())
	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestLoggerDebug(t *testing.T) {
	msg := "debug"
	testLogger.Debug(msg)
	zaplogger.Debug(msg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerDebugf(t *testing.T) {
	msg := "debugf %s"
	arg := "foo"
	testLogger.Debugf(msg, arg)
	zaplogger.Debugf(msg, arg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerDebugln(t *testing.T) {
	msg := "debugln"
	testLogger.Debugln(msg)
	zaplogger.Debug(msg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerInfo(t *testing.T) {
	msg := "info"
	testLogger.Info(msg)
	zaplogger.Info(msg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerInfof(t *testing.T) {
	msg := "infof %s"
	arg := "foo"
	testLogger.Infof(msg, arg)
	zaplogger.Infof(msg, arg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerInfoln(t *testing.T) {
	msg := "infoln"
	testLogger.Infoln(msg)
	zaplogger.Info(msg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerWarn(t *testing.T) {
	msg := "warn"
	testLogger.Warn(msg)
	zaplogger.Warn(msg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerWarnf(t *testing.T) {
	msg := "warnf %s"
	arg := "foo"
	testLogger.Warnf(msg, arg)
	zaplogger.Warnf(msg, arg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerWarnln(t *testing.T) {
	msg := "warnln"
	testLogger.Warnln(msg)
	zaplogger.Warn(msg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerError(t *testing.T) {
	msg := "error"
	testLogger.Error(msg)
	zaplogger.Error(msg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerErrorf(t *testing.T) {
	msg := "errorf %s"
	arg := "foo"
	testLogger.Errorf(msg, arg)
	zaplogger.Errorf(msg, arg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func TestLoggerErrorln(t *testing.T) {
	msg := "errorln"
	testLogger.Errorln(msg)
	zaplogger.Error(msg)
	if err := validateLoggerOutput(); err != nil {
		t.Error(err)
	}
}

func validateLoggerOutput() error {
	actual, err := ioutil.ReadAll(tempLogFile1)
	if err != nil {
		return err
	}
	expected, err := ioutil.ReadAll(tempLogFile2)
	if err != nil {
		return err
	}
	if !bytes.Equal(actual, expected) {
		return fmt.Errorf("expected %s, got %s", expected, actual)
	}

	return nil
}

func ExampleNew() {
	zapLogger, err := zap.NewDevelopment(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := New(zapLogger.Sugar())
	logger.Debug("debug")
	logger.Debugf("debugf")
	logger.Debugln("debugln")

	logger.Info("info")
	logger.Infof("infof")
	logger.Infoln("infoln")

	logger.Warn("warn")
	logger.Warnf("warnf")
	logger.Warnln("warnln")

	logger.Error("error")
	logger.Errorf("errorf")
	logger.Errorln("errorln")

	//logger.Fatal("fatal")
	//logger.Fatalf("fatalf")
	//logger.Fatalln("fatalln")
}
