package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	chars := make(map[rune]struct{})
	runes := []rune(strings.ToLower(s))
	for _, rn := range runes {
		_, ok := chars[rn]
		if !ok {
			chars[rn] = struct{}{}
		} else {
			return false
		}

	}
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Printf("'%s' - %t\n", s1, isUnique(s1))
	fmt.Printf("'%s' - %t\n", s2, isUnique(s2))
	fmt.Printf("'%s' - %t\n", s3, isUnique(s3))
}
