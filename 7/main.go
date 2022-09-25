package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func Writer(data map[int]string, mx *sync.Mutex, ctx context.Context) {
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
			key := rand.Intn(10)
			val := make([]byte, 5)
			rand.Read(val)
			mx.Lock()
			data[key] = string(val)
			mx.Unlock()
			/*ну короче просто мьютексом мапку закрываем, так что единовременно к ней обращается одна горутина. можно юзать syncMap,
			в которую эти мьютексы встроены
			 */
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	mx := sync.Mutex{}
	data := map[int]string{}
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	for i := 0; i < num; i++ {
		go Writer(data, &mx, ctx)
	}
	time.Sleep(time.Second)
	cancel()
	for key, val := range data {
		fmt.Println(key, ": ", val)
	}
}
