package main

import (
	"context"
	"fmt"
	"math/rand"

	// "math/rand"
	"sync"
	"time"
)

// Суть алгоритма:
// 1) У пользователя есть "корзина" с токенами.
// 2) В корзине может быть максимум maxTokens (например, 5).
// 3) Каждые N миллисекунд в корзину добавляется 1 токен (регенерация).
// 4) Когда приходит запрос, вызывается метод Allow().
// 5) Если в корзине есть хотя бы один токен — мы его забираем и возвращаем true (запрос разрешен).
// 6 )Если токенов нет — возвращаем false (запрос отклонен).

// Задача:
// Реализовать структуру RateLimiter, которая будет работать по этому принципу. Она должна быть потокобезопасной и эффективной.

// 1) Thread-safety: Корректная работа при одновременных вызовах Allow().
// 2) Lazy Refill (Важно!): Вместо того чтобы запускать отдельную горутину с тикером, которая будет добавлять токены в фоне (это ест ресурсы), попробуй реализовать "ленивое" пополнение. То есть вычислять, сколько токенов должно было добавиться, прямо в момент вызова Allow(), исходя из прошедшего времени.
// 3) Точность: Правильный расчет интервалов времени.


type RateLimiter struct {
	maxTokens float64
	refillRate float64
	tokens float64
	lastRefillTime time.Time
	mu sync.Mutex
}

func NewRateLimiter(maxTokens int, refillInterval time.Duration) *RateLimiter {
	newRateLimiter := &RateLimiter{
		maxTokens: float64(maxTokens),
		refillRate: 1 / float64(refillInterval.Nanoseconds()),
		tokens: float64(maxTokens),
		lastRefillTime: time.Now(),
		mu: sync.Mutex{},
	}

    return newRateLimiter
}

func (rl *RateLimiter) Allow(ctx context.Context) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefillTime).Nanoseconds()

	rl.tokens += float64(elapsed) * rl.refillRate

	if rl.tokens > rl.maxTokens {
		rl.tokens = rl.maxTokens
	}

	rl.lastRefillTime = now

	if rl.tokens >= 1.0 {
		rl.tokens--
		return true
	}

	return false
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	rl := NewRateLimiter(10, 500 * time.Millisecond)
	wg := sync.WaitGroup{}

	wg.Add(30)
	for i := 0; i < 30; i++ {
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		go func() {
			defer wg.Done()
			fmt.Println(rl.Allow(ctx))
		}()
	}

	wg.Wait()
}