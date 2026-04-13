package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value string
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {return len(pq)}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func main() {
	items := map[string]int{
		"Задача 1": 3, "Задача 2": 1, "Задача 3": 5,
	}
	
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value: value,
			priority: priority,
			index: i,
		}
		i++
	}
	heap.Init(&pq)
	
	newItem := &Item{value: "Задача 4", priority: 4}
	heap.Push(&pq, newItem)
	
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s\n", item.priority, item.value)
	}
}