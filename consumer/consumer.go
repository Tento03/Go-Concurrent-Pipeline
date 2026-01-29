package consumer

import (
	"concurrent-pipeline/metrics"
	"concurrent-pipeline/models"
	"context"
)

func StartConsumer(ctx context.Context, in <-chan models.Result, metrics *metrics.Metrics) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case res, ok := <-in:
				if !ok {
					return
				}

				if res.Success {
					metrics.IncProcessed()
				} else {
					metrics.IncFailed()
				}
			}
		}
	}()
}
