package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

type Action struct {
	Human
	Job string
}

func (a *Action) ShowJob() {
	fmt.Printf("%s %s's job is %s.", a.Name, a.Surname, a.Job)
}

func (h *Human) Greet() {
	fmt.Printf("Hello, %s %s!\n", h.Name, h.Surname)
}

func (h *Human) ShowAge() {
	fmt.Printf("%s %s's age is %d.\n", h.Name, h.Surname, h.Age)
}

func main() {
	h := Human{
		Name:    "Kirill",
		Surname: "Kozhushnyy",
		Age:     22,
	}
	a := Action{
		Human: h,
		Job:   "hobo",
	}
	a.Greet()
	a.ShowAge()
	a.ShowJob()
}
