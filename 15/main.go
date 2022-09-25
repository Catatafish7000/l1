package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	ans := strings.Repeat("a", size)
	return ans
}
func someFunc() string {
	v := createHugeString(1 << 10)
	//justString = v[:100]
	justString := make([]rune, 100)
	copy(justString, []rune(v)[:100])
	return string(justString)
}

//В оригинальном коде итерация идёт по байтам, а не рунам,что приведёт к броблемам с символами крупнее одного байта

func main() {
	fmt.Println(someFunc())
}
