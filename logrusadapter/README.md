# logrus logger adapter

## Usage

```go
import (
    "github.com/clevergo/log"
    "github.com/clevergo/log/logrusadapter"
    "github.com/sirupsen/logrus"
)

var logger log.Logger

func main() {
    logger = logrusadapter.New(logrus.New())
    logger.Debug("debug msg")
    // ...
}
```
