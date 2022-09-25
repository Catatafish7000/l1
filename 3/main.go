package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	result := 0
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			mx.Lock()
			result += val * val
			mx.Unlock()
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println(result)
}
