package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
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
Loop1:
	for {
		select {
		case <-ctx.Done():
			break Loop1
		case data := <-ch:
			fmt.Println(string(data))
		}
	}
}

func main() {
	var secs int
	var _, err = fmt.Scanf("%d", &secs)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	wt := time.Duration(secs) * time.Second
	ch := make(chan []byte)
	defer close(ch)
	ctx, _ := context.WithTimeout(context.Background(), wt)
	go Reader(ctx, ch)
	go Writer(ctx, ch)
	time.Sleep(wt)
	fmt.Println("Done")
}