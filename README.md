# Go logging rethinked

[![Circle CI](https://circleci.com/gh/yanzay/log.svg?style=svg)](https://circleci.com/gh/yanzay/log)
[![Build Status](https://travis-ci.org/yanzay/log.svg?branch=master)](https://travis-ci.org/yanzay/log)
[![Coverage Status](https://coveralls.io/repos/github/yanzay/log/badge.svg?branch=master)](https://coveralls.io/github/yanzay/log?branch=master)

Usage:

```
log.Println("some log") // unconditional log
log.Trace("trace") // log only with `trace` level
log.Tracef("42: %s", "yep") // each method has it's format alternative
log.Debug("debug") // log only with `debug` level and lower
log.Info("info") // log only with `info` level and lower
log.Warning("warn") // log with `warning` level and lower
log.Error("err") // log with `error` and `critical` level
log.Fatal("haha") // log and panic("haha")
```

Log adds `--log-level` flag to your program:

```
// main.go
package main

import (
    "flag"

    "github.com/yanzay/log"
)

func main() {
    flag.Parse()
    log.Info("info")
}
```

```
$ go run main.go --help
Usage:
  -log-level string
        Log level: trace|debug|info|warning|error|critical (default "info")
```

## Advanced Usage

You can set logging level manually by:
```
log.Level = log.LevelTrace
```

Also you can use your own log writer:

```
log.Writer = myWriter // myWriter should implement io.Writer interface
```

