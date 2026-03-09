package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// BenchmarkAtomicCounter-12    	    1538	    766267 ns/op	  120333 B/op	    5002 allocs/op
// PASS
// ok  	github.com/stepan41k/GolangTasks/cmd	1.476s
func GoroutineAtomicCounter(a int) int64 {
	var counter atomic.Int64
	wg := sync.WaitGroup{}

	wg.Add(a)
	for i := 0; i < a; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()

	return counter.Load()
}

// BenchmarkMutexCounter-12    	    1556	    767217 ns/op	  160601 B/op	    5003 allocs/op
// PASS
// ok  	github.com/stepan41k/GolangTasks/cmd	1.480s
func GoroutineMutexCounter(a int) int {
	counter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(a)
	for i := 0; i < a; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	return counter
}

func main() {
	// val := GoroutineCounterAtomic(5000)
	val := GoroutineMutexCounter(5000)
	fmt.Println(val)
}
