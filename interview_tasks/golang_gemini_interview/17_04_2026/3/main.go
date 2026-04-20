package main

import (
	"fmt"
	"sync"
	"time"
)

func StartPipeline(in <-chan int, out chan<- int, rateLimit int) {
	interval := time.Second / time.Duration(rateLimit)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	var wg sync.WaitGroup
	
	for val := range in {
		<-ticker.C
		
		wg.Add(1)
		
		go func(v int) {
			defer wg.Done()
			
			time.Sleep(100 * time.Millisecond)
			
			out <- v
		}(val)
	}
	
	wg.Wait()
	close(out)
}

func main() {
	inChan := make(chan int)
	outChan := make(chan int)
	limit := 5

	go func() {
		for i := range 100 {
			inChan <- i
		}

		close(inChan)
	}()

	go func() {
		StartPipeline(inChan, outChan, limit)
	}()

	for x := range outChan {
		fmt.Println(x)
	}
}
