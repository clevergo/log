// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package zapadapter

import (
	"testing"

	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	zapLogger, _ := zap.NewDevelopment(zap.AddCallerSkip(1))
	_ = New(zapLogger.Sugar())
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
