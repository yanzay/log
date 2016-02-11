# Go logging rethinked

Usage:

```
log.Println("some log") // unconditional log
log.Trace("trace") // log only with `trace` level
log.Tracef("42: %s", "yep") // each method has it's format alternative
log.Debug("debug") // log only with `debug` level and lower
log.Info("info") // log only with `info` level and lower
log.Warning("warn") // log with `warning` level and lower
log.Error("err") // log with `error` and `critical` level
log.Fatal("haha") // log and os.Exit(1)
```

Log adds `--log-level` flag to your program:

```
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
