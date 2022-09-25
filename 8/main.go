package main

import (
	"fmt"
	"log"
)

func main() {
	var i int
	var val int64
	fmt.Scanf("%d %d", &i, &val)
	l := len(fmt.Sprintf("%b", val))
	if i > l {
		log.Fatalf("%s", "Bad index")
	}
	changer := int64(1 << int64(len(fmt.Sprintf("%b", val))-i-1)) //все биты, кроме i-го, по нулям
	fmt.Printf("Ставим на 0: %d\n", val&^changer)                 //побитовое  не и
	fmt.Printf("Ставим на 1: %d\n", val|changer)                  //побитовое или
}
