package main

import (
	"sync"
	"time"
)

// Контекст:
// Нам нужно реализовать кэш в оперативной памяти. Кэш — это структура данных, которая хранит значения по ключам. Однако память не бесконечна, поэтому у каждой записи есть "срок годности" (TTL). По истечении этого срока запись должна считаться невалидной и удаляться.

// Твоя задача:
// 1) Реализовать структуру TTLCache, которая:
// 2) Имеет метод Set(key string, value interface{}, ttl time.Duration).
// 3) Имеет метод Get(key string) (interface{}, bool).
// 4) Поддерживает Thread-safety (безопасную работу из разных горутин).
// 5) Оптимизирована для чтения: условие задачи — чтений (Get) в 100 раз больше, чем записей (Set).
// 6) Имеет механизм автоматической очистки просроченных данных (чтобы память не утекала).

type item struct {
	value       interface{}
	expiresAt   time.Time
}

type TTLCache struct {
	items map[string]item
	mx sync.RWMutex
	cleanupInterval time.Duration
	done chan struct{}
}

func NewTTLCache(cleanupInterval time.Duration) *TTLCache {
	ticker := time.NewTicker(cleanupInterval)
	
	newCache := &TTLCache{
		items: make(map[string]item),
		mx: sync.RWMutex{},
		cleanupInterval: cleanupInterval,
		done: make(chan struct{}),
	}

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				newCache.cleanup()
			case <-newCache.done:
				return
			}
		}
	}()

    return newCache
}


func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mx.Lock()
	c.items[key] = item{value: value, expiresAt: time.Now().Add(ttl)}
	c.mx.Unlock()
}

func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.mx.RLock()
	val, ok := c.items[key]
	c.mx.RUnlock()

	if !ok  {
		return nil, false
	}

	if val.expiresAt.Before(time.Now()) {
		c.mx.Lock()
		if val, ok := c.items[key]; ok && val.expiresAt.Before(time.Now()) {
			delete(c.items, key)
		}
		c.mx.Unlock()
		return nil, false
	} 

	return val.value, true
}

func (c *TTLCache) cleanup() {
	c.mx.Lock()
	defer c.mx.Unlock()
	
	startTime := time.Now()

    for key, val := range c.items {
		if time.Since(startTime) > 1 * time.Second {
			return
		}

		if val.expiresAt.Before(startTime) {
			delete(c.items, key)
		} 
	}
}