package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать структуру Batcher, которая:
// Принимает строки (события) через метод Add.
// "Сбрасывает" (flush) батч, если:
// Набралось maxBatchSize элементов.
// Прошло flushInterval времени с момента последнего сброса (таймаут).
// Должна быть потокобезопасной.
// Должна корректно завершать работу (graceful shutdown) при отмене контекста, сбрасывая оставшиеся в памяти данные.

type Batcher struct {
	maxBatchSize int
	flushInterval time.Duration
	inputCh chan string
	saveFunc func([]string)
}

func NewBatcher(maxBatchSize int, flushInterval time.Duration, saveFn func([]string)) *Batcher {
    return &Batcher{
		maxBatchSize: maxBatchSize,
		flushInterval: flushInterval,
		inputCh: make(chan string, maxBatchSize),
		saveFunc: saveFn,
	}
}

func (b *Batcher) Add(ctx context.Context, event string) error {
	select {
	case b.inputCh <- event:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
	
}

func (b *Batcher) Run(ctx context.Context) {
	batch := make([]string, 0, b.maxBatchSize)
	ticker := time.NewTicker(b.flushInterval)
	defer func() {
		ticker.Stop()
	}()

	flush := func() {
		if len(batch) > 0 {
			res := make([]string, len(batch))
			copy(res, batch)
			b.saveFunc(res)
			batch = batch[:0]
		}
	}

	for {
		select {
		case <-ctx.Done():
			for {
				select {
				case item := <-b.inputCh:
					batch = append(batch, item)
					if len(batch) >= b.maxBatchSize {
						flush()
					}
				default:
					flush()
					return
				}
			}
		case item := <-b.inputCh:
			batch = append(batch, item)
			if len(batch) >= b.maxBatchSize {
				flush()
				ticker.Reset(b.flushInterval)
			}
		case <-ticker.C:
			flush()
		}
	}
}

func main() {
	wg := sync.WaitGroup{}

	save := func(batch []string) {
		fmt.Println("Saved batch of size:", len(batch), batch)
	}

	b := NewBatcher(5, 2*time.Second, save)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		b.Run(ctx)
	}()

	b.Add(ctx, "event 1")
	b.Add(ctx, "event 2")
	b.Add(ctx, "event 3")
	b.Add(ctx, "event 4")
	b.Add(ctx, "event 5")
	b.Add(ctx, "event 6")
	b.Add(ctx, "event 7")

	wg.Wait()
}