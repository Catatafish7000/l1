package main

import "fmt"

var strings = []string{"cat", "cat", "dog", "cat", "tree"}

func main() {
	set := make(map[string]struct{})

	for _, str := range strings {
		set[str] = struct{}{}
	}
	fmt.Println(set)
}
