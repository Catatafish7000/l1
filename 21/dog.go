package main

import "fmt"

type dog interface {
	Bark()
}
type hound struct {
}

func (h *hound) Bark() {
	fmt.Println("Гав")
}
