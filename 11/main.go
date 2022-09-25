package main

import "fmt"

func GetIntersec(a []int, b []int) []int {
	ans := []int{}
	for _, v_a := range a {
		for _, v_b := range b {
			if v_a == v_b {
				ans = append(ans, v_b)
			}
		}
	}
	return ans
}

func main() {
	mn11 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mn12 := []int{2, 5, 6, 9, 15, 18}
	fmt.Println(GetIntersec(mn12, mn11))
}
