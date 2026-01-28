package worker

import (
	"concurrent-pipeline/models"
	"context"
	"sync"
	"time"
)

func StartWorkerPool(ctx context.Context, workerCount int, in <-chan models.Data, out chan<- models.Result) {
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-in:
					if !ok {
						return
					}

					time.Sleep(300 * time.Millisecond)

					out <- models.Result{
						DataID:  data.ID,
						Success: true,
					}
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}
