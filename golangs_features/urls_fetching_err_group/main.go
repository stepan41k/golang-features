package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func fetchURL(ctx context.Context, url string) (string, error) {
	select {
	case <-time.After(time.Duration(100+len(url)*10) * time.Millisecond):
		return "data from " + url, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func FetchAll(ctx context.Context, urls []string) ([]string, error) {
	eg, ctx := errgroup.WithContext(ctx)
	results := make([]string, len(urls))
	semaphore := make(chan struct{}, 3)

	for ind, url := range urls {
		select {
		case semaphore <- struct{}{}:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
		
		ind, url := ind, url

		eg.Go(func() error {
			defer func() {
				<-semaphore
			}()
			
			result, err := fetchURL(ctx, url)
			if err != nil {
				return err
			}

			results[ind] = result
			return nil
		})
	}

	close(semaphore)

	if err := eg.Wait(); err != nil {
		fmt.Printf("expected error: %s", err.Error())
		return nil, fmt.Errorf("founded error: %w", err)
	}

    return results, nil
}

func main() {
	urls := []string{"google.com", "yandex.ru", "github.com", "stackoverflow.com", "go.dev"}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := FetchAll(ctx, urls)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Results:", data)
}