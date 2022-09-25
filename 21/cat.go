package main

import "fmt"

type cat interface {
	Meow()
}

type kitty struct{}

func (w *kitty) Meow() {
	fmt.Println("Мяу")
}
