package main

import (
	"fmt"
	"time"
)

const (
	MaxNumbers = 9999
)

type Consumer interface {
	Consume(chunk []any) (err error)
}

type Producer interface {
	Next() (chunk []any, cookie int, err error)
	Commit(cookie int) (err error)
}

func Ping(p Producer, c Consumer) error {
	batch := make([]any, 0, MaxNumbers)
	var pendingCookies []int 
	
	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	flush := func() error {
		if len(batch) == 0 {
			return nil
		}

		if err := c.Consume(batch); err != nil {
			return fmt.Errorf("consume error: %w", err)
		}

		for _, cookie := range pendingCookies {
			if err := p.Commit(cookie); err != nil {
				return fmt.Errorf("commit error (cookie %d): %w", cookie, err)
			}
		}

		batch = batch[:0]
		pendingCookies = pendingCookies[:0]
		return nil
	}

	for {
		chunk, cookie, err := p.Next()
		if err != nil {
			return fmt.Errorf("producing error: %w", err)
		}

		if len(batch)+len(chunk) > MaxNumbers {
			if err := flush(); err != nil {
				return err
			}
		}

		batch = append(batch, chunk...)
		pendingCookies = append(pendingCookies, cookie)

		select {
		case <-ticker.C:
			if err := flush(); err != nil {
				return err
			}
		default:
		}
	}
}

func main() {

}