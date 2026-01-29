package logger

import (
	"concurrent-pipeline/metrics"
	"context"
	"fmt"
	"time"
)

func StartLogger(ctx context.Context, metrics *metrics.Metrics, duration time.Duration) {
	go func() {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				p, f := metrics.Snapshot()
				fmt.Printf("Processed:%d | Failed:%d\n", p, f)
			}
		}
	}()
}
