package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
)

func Writer(ctx context.Context, ch chan<- []byte) {
Loop:
	for {
		data := make([]byte, 10)
		rand.Read(data)
		select {
		case <-ctx.Done():
			break Loop
		default:
			ch <- data
		}
	}
}

func Reader(ctx context.Context, ch <-chan []byte) {
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		case data := <-ch:
			fmt.Println(string(data))
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var num int
	var _, err = fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	ch := make(chan []byte)
	defer close(ch)
	for i := 0; i < num; i++ {
		go Reader(ctx, ch)
	}
	go Writer(ctx, ch)
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
	cancel()
}
