# Generic leveled logger interface
[![Build Status](https://travis-ci.org/clevergo/log.svg?branch=master)](https://travis-ci.org/clevergo/log)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/clevergo.tech/log?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/log)](https://goreportcard.com/report/github.com/clevergo/log)
[![Release](https://img.shields.io/github/release/clevergo/log.svg?style=flat-square)](https://github.com/clevergo/log/releases)
[![Sourcegraph](https://sourcegraph.com/github.com/clevergo/log/-/badge.svg)](https://sourcegraph.com/github.com/clevergo/log?badge)

## Usage

Checkout [example](https://github.com/clevergo/examples/tree/master/log) for details.

### Zap

```go
var logger log.Logger = logrus.New()
```

### Logrus

```go
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
