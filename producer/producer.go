package producer

import (
	"concurrent-pipeline/models"
	"context"
	"time"
)

func StartProducer(ctx context.Context, out chan<- models.Data) {
	go func() {
		defer close(out)

		id := 0
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				id++
				out <- models.Data{
					ID:        id,
					Payload:   "data",
					CreatedAt: t,
				}
			}
		}
	}()
}
