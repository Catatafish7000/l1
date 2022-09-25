package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Через контекст:")
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping1")
				return
			default:
				data := make([]byte, 10)
				rand.Read(data)
				fmt.Println(string(data))
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	fmt.Println("\n\nЧерез канал:")
	ch := make(chan bool)
	defer close(ch)
	go func(closing <-chan bool) {
		for {
			select {
			case <-closing:
				fmt.Println("Stopping2")
				return
			default:
				data := make([]byte, 10)
				rand.Read(data)
				fmt.Println(string(data))
			}
		}
	}(ch)
	time.Sleep(10 * time.Second)
	ch <- true
}
