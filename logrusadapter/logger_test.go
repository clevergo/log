// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package logrusadapter

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	_ = New(logrus.New())
}

func ExampleNew() {
	logger := New(logrus.New())
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
