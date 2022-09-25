package main

import (
	"fmt"
	"time"
)

func Sleep(dur int) {
	start := time.Now()
	end := start.Add(time.Second * time.Duration(dur))
	for time.Now().UnixNano() < end.UnixNano() {
	}
}

func main() {
	start := time.Now()
	Sleep(15)
	end := time.Now()
	dur := end.Sub(start)
	fmt.Println(dur)
}
