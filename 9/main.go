package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer close(ch1)
	defer close(ch2)
	defer fmt.Println("Done")
	ctx, cancel := context.WithCancel(context.Background())
	go func(ch chan<- int, ctx context.Context) {
	Loop1:
		for {
			select {
			case <-ctx.Done():
				break Loop1
			default:
				val := rand.Intn(10)
				ch <- val
			}
		}
	}(ch1, ctx)
	go func(ch1 <-chan int, ch2 chan<- int, ctx context.Context) {
	Loop2:
		for {
			select {
			case <-ctx.Done():
				break Loop2
			default:
				val := <-ch1
				ch2 <- val * val
			}
		}
	}(ch1, ch2, ctx)
	go func(ch2 <-chan int) {
		for val := range ch2 {
			fmt.Println(val)
		}
	}(ch2)
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
	cancel()
}
