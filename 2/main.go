package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			fmt.Println(val * val)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
