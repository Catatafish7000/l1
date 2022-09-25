package main

func Voice(d dog) {
	d.Bark()
}
func main() {
	cat := &kitty{}
	dog := &hound{}
	adapter := &catAdapter{
		cat: cat,
	}
	Voice(dog)
	Voice(adapter)
}
