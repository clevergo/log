# Generic leveled logger interface
[![Build Status](https://travis-ci.org/clevergo/log.svg?branch=master)](https://travis-ci.org/clevergo/log)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/log)](https://goreportcard.com/report/github.com/clevergo/log)
[![Release](https://img.shields.io/github/release/clevergo/log.svg?style=flat-square)](https://github.com/clevergo/log/releases)
[![Sourcegraph](https://sourcegraph.com/github.com/clevergo/log/-/badge.svg)](https://sourcegraph.com/github.com/clevergo/log?badge)

## Usage

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

### Example

```shell
$ cd example

$ go run main.go -adapter zap   
{"level":"debug","msg":"Debug"}
{"level":"debug","msg":"Debugf"}
{"level":"info","msg":"Info"}
{"level":"info","msg":"Infof"}
{"level":"warn","msg":"Warn"}
{"level":"warn","msg":"Warnf"}
{"level":"error","msg":"Error"}
{"level":"error","msg":"Errorf"}

$ go run main.go -adapter logrus 
INFO[0000] Info                                         
INFO[0000] Infof                                        
WARN[0000] Warn                                         
WARN[0000] Warnf                                        
ERRO[0000] Error                                        
ERRO[0000] Errorf   
```
