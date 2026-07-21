package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type Result struct {
	URL string
	StatusCode int
}


func fetchUrls(urls []string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.SetLimit(3)

	mu := sync.Mutex{}
	results := []Result{}

	for _, url := range urls {
		eg.Go(func() error {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return fmt.Errorf("failed to create request for %s: %w", url, err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				return fmt.Errorf("failed fetching %s: %w", url, err)
			}

			defer resp.Body.Close()

			resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("url %s returned status %d", url, resp.StatusCode)
			}

			mu.Lock()
			results = append(results, Result{URL: url, StatusCode: resp.StatusCode})
			mu.Unlock()

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}


func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://bad-url-that-fails.com",
		"https://go.dev",
		"https://yandex.ru",
	}

	results, err := fetchUrls(urls)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}

	fmt.Println(results)
}

package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// Graceful Shutdown / Ошибки: Если один из запросов вернул критическую ошибку, нужно отменить все остальные текущие запросы и вернуть ошибку немедленно.

type User struct {
	ID   int
	Name string
}

func fetchUser(ctx context.Context, id int) (*User, error) {
	select {
	case <-time.After(time.Millisecond * 500):
		return &User{ID: id, Name: fmt.Sprintf("User-%d", id)}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	ids := make([]int, 100)
	for i := range ids {
		ids[i] = i + 1
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	users, err := ProcessUsers(ctx, ids, 10)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Successfully fetched %d users\n", len(users))
}

func ProcessUsers(ctx context.Context, ids []int, concurrency int) ([]*User, error) {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	res := make([]*User, len(ids))

	for idx, id := range ids {
		id := id

		g.Go(func() error {
			user, err := fetchUser(ctx, id)
			if err != nil {
				return err
			}

			res[idx] = user
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}
