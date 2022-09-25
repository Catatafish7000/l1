package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	count int64
}

func (c *counter) add(val int64) {
	atomic.AddInt64(&c.count, val)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	cnt := counter{count: 0}
	var mx sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				mx.Lock()
				cnt.count += 1
				mx.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt.count) //через мьютекс
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				cnt.add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt.count) //ещё 100 с помощью атомарной операции
}
