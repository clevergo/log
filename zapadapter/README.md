# zap logger adapter

## Usage

```go
import (
    "github.com/clevergo/log"
    "github.com/clevergo/log/zapadapter"
    "go.uber.org/zap"
)

var logger log.Logger

func main() {
    zapLogger, err := zap.NewDevelopment(zap.AddCallerSkip(1))
    if err != nil {
        panic(err)
    }
    logger := New(zapLogger.Sugar())
    logger.Debug("debug msg")
    // ...
}
```
