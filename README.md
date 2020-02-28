# Universal leveled logger interface.
[![Build Status](https://travis-ci.org/clevergo/log.svg?branch=master)](https://travis-ci.org/clevergo/log) 
[![Coverage Status](https://coveralls.io/repos/github/clevergo/log/badge.svg?branch=master)](https://coveralls.io/github/clevergo/log?branch=master) 
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/log) 
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/log)](https://goreportcard.com/report/github.com/clevergo/log) 
[![Release](https://img.shields.io/github/release/clevergo/log.svg?style=flat-square)](https://github.com/clevergo/log/releases)

## Adapters

- [logrus](https://github.com/clevergo/log-logrus)
- [zap](https://github.com/clevergo/log-zap)

## Usage

```go
logger.Debug(args ...interface{})
logger.Debugf(format string, args ...interface{})
logger.Debugln(args ...interface{})

logger.Fatal(args ...interface{})
logger.Fatalf(format string, args ...interface{})
logger.Fatalln(args ...interface{})

logger.Info(args ...interface{})
logger.Infof(format string, args ...interface{})
logger.Infoln(args ...interface{})

logger.Warn(args ...interface{})
logger.Warnf(format string, args ...interface{})
logger.Warnln(args ...interface{})

logger.Error(args ...interface{})
logger.Errorf(format string, args ...interface{})
logger.Errorln(args ...interface{})
```
