# `github.com/boostgo/log`

# Get started

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/boostgo/log"
	"github.com/boostgo/log/logx"
	"github.com/boostgo/trace"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	logx.Pretty()
	logx.SetCancel(cancel)
	logx.InitTracer(trace.NewTracer().IAmMaster(true))

	log.Info().Msg("info msg")   // 4:53PM INF info msg
	log.Fatal().Msg("fatal msg") // 4:53PM ERR fatal msg fatal=true

	select {
	case <-ctx.Done():
		fmt.Println("context canceled") // context canceled
		os.Exit(1)
	default:
	}

	log.Debug().Msg("debug msg") // does not print because called fatal with context cancel
}

```
