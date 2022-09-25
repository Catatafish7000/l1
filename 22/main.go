package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewInt(int64(2 * math.Pow(10, 20)))
	b := big.NewInt(int64(2 * math.Pow(10, 20)))

	mul := big.NewInt(0).Mul(a, b)
	div := big.NewInt(0).Div(a, b)
	sum := big.NewInt(0).Add(a, b)
	diff := big.NewInt(0).Sub(a, b)
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("x * y = %d\n", mul)
	fmt.Printf("x / y = %d\n", div)
	fmt.Printf("x + y = %d\n", sum)
	fmt.Printf("x - y = %d\n", diff)
}
