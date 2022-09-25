package main

type catAdapter struct {
	cat *kitty
}

func (c *catAdapter) Bark() {
	c.cat.Meow()
}
