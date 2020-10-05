# Generic leveled logger interface
[![Build Status](https://img.shields.io/travis/clevergo/log?style=flat-square)](https://travis-ci.org/clevergo/log)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/log?style=flat-square)](https://coveralls.io/github/clevergo/log)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/clevergo.tech/log?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/log?style=flat-square)](https://goreportcard.com/report/github.com/clevergo/log)
[![Release](https://img.shields.io/github/release/clevergo/log.svg?flat-square)](https://github.com/clevergo/log/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/total/clevergo.tech/log&style=flat-square)](https://pkg.clevergo.tech/)
[![Chat](https://img.shields.io/badge/chat-telegram-blue?style=flat-square)](https://t.me/clevergotech)
[![Community](https://img.shields.io/badge/community-forum-blue?style=flat-square&color=orange)](https://forum.clevergo.tech)

## Usage

Checkout [example](https://github.com/clevergo/examples/tree/master/log) for details.

```shell
go get -u clevergo.tech/log
```

### Standard Logger

`StdLogger` wraps Go standard logger `log.Logger`.

```go
import (
    stdlog "log"

    "clevergo.tech/log"
)

var logger log.Logger = log.New(os.Stderr, "", stdlog.LstdFlags)
```

### Logrus

```go
import (
    "clevergo.tech/log"
    "github.com/sirupsen/logrus"
)

var logger log.Logger = logrus.New()
```

### Zap

```go
import (
    "clevergo.tech/log"
    "go.uber.org/zap"
)

var logger log.Logger = zap.NewExample().Sugar()
```

### Interface

```go
logger.Debug(args ...interface{})
logger.Debugf(format string, args ...interface{})
logger.Info(args ...interface{})
logger.Infof(format string, args ...interface{})
logger.Warn(args ...interface{})
logger.Warnf(format string, args ...interface{})
logger.Error(args ...interface{})
logger.Errorf(format string, args ...interface{})
logger.Fatal(args ...interface{})
logger.Fatalf(format string, args ...interface{})
```
