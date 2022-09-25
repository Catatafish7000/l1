package main

import "fmt"

func main() {
	a := 24
	b := 37
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)
}
