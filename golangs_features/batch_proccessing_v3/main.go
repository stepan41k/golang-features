package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func fetchBatch(ids []int) (map[int]string, error) {
	fmt.Printf("[DB] Fetching batch: %v\n", ids)
	results := make(map[int]string)
	for _, id := range ids {
		results[id] = fmt.Sprintf("Result_for_%d", id)
	}

	time.Sleep(4 * time.Millisecond)
	return results, nil
}

type response struct {
	data string
	err  error
}

type request struct {
	id       int
	respChan chan response
}

type BatchLoader struct {
	requestChan  chan request
	linger       time.Duration
	maxBatchSize int
	firstInsert  time.Time
}

func NewBatchLoader(maxBatchSize int, linger time.Duration) *BatchLoader {
	batchLoader := &BatchLoader{
		requestChan:  make(chan request, maxBatchSize),
		linger:       linger,
		maxBatchSize: maxBatchSize,
	}

	go func() {
		batchLoader.start()
	}()

	return batchLoader
}

func (l *BatchLoader) Load(ctx context.Context, id int) (string, error) {
	respCh := make(chan response, 1)

	select {
	case l.requestChan <- request{id: id, respChan: respCh}:
	case <-ctx.Done():
		return "", ctx.Err()
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case resp := <-respCh:
		return resp.data, resp.err
	}
}

func (l *BatchLoader) start() {
	ticker := time.NewTicker(l.linger)
	defer ticker.Stop()
	requestBatch := make([]request, 0, l.maxBatchSize)

	flush := func() {
		if len(requestBatch) > 0 {
			copyBatch := make([]int, 0, len(requestBatch))

			for _, v := range requestBatch {
				copyBatch = append(copyBatch, v.id)
			}

			results, err := fetchBatch(copyBatch)
			for _, req := range requestBatch {
				if err != nil {
					req.respChan <- response{data: "", err: err}
					continue
				}

				val, ok := results[req.id]
				if !ok {
					req.respChan <- response{data: "", err: fmt.Errorf("not found")}
					continue
				}

				req.respChan <- response{data: val, err: nil}
			}

			requestBatch = requestBatch[:0]
		}
	}

	for {
		select {
		case <-ticker.C:
			flush()
		case val, ok := <-l.requestChan:
			if !ok {
				flush()
				return
			}
			requestBatch = append(requestBatch, val)
			if len(requestBatch) >= l.maxBatchSize {
				flush()
				ticker.Reset(l.linger)
			}
		}
	}
}

func main() {
	loader := NewBatchLoader(5, 20*time.Millisecond)
	wg := sync.WaitGroup{}

	for i := 1; i <= 12; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			res, err := loader.Load(ctx, id)
			if err != nil {
				fmt.Printf("ID %d Error: %v\n", id, err)
			} else {
				fmt.Printf("ID %d Success: %s\n", id, res)
			}
		}(i)
	}

	wg.Wait()
}
