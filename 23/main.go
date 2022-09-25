package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	sl := make([]int, 100, 100)
	for i, _ := range sl {
		sl[i] = rand.Intn(200)
	}
	fmt.Println(sl)
	var idx int
	fmt.Scanf("%d", &idx)
	if idx > len(sl)-1 {
		log.Fatalf("out of slice range")
	}
	if idx == len(sl)-1 {
		sl = sl[:len(sl)-1]
	} else {
		sl = append(sl[:idx], sl[idx+1:]...)
	}
	fmt.Println(sl)
}
