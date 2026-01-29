package main

import (
	"concurrent-pipeline/consumer"
	"concurrent-pipeline/logger"
	"concurrent-pipeline/metrics"
	"concurrent-pipeline/models"
	"concurrent-pipeline/producer"
	"concurrent-pipeline/worker"
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan models.Data)
	out := make(chan models.Result)

	metrics := &metrics.Metrics{}

	go producer.StartProducer(ctx, in)

	worker.StartWorkerPool(ctx, 4, in, out)

	consumer.StartConsumer(ctx, out, metrics)

	logger.StartLogger(ctx, metrics, 2*time.Second)

	time.Sleep(10 * time.Second)
}
