package main

import "fmt"

func main() {
	i := 12
	s := "12"
	b := true
	ch := make(chan int)
	fl := 1.25
	fu := func() {}
	an := []any{i, s, b, ch, fl, fu}
	for _, v := range an {
		fmt.Println(v, fmt.Sprintf("%T", v))
	}
}
