package main

import (
	"fmt"
	"math/rand"
)

func qSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	} else {
		pivot := slice[0]

		less := make([]int, 0)
		greater := make([]int, 0)

		for _, val := range slice[1:] {
			if val < pivot {
				less = append(less, val)
			} else {
				greater = append(greater, val)
			}
		}

		less = append(qSort(less), pivot)
		greater = qSort(greater)
		return append(less, greater...)
	}
}

func main() {
	arr := make([]int, 100, 100)
	for i, _ := range arr {
		arr[i] = rand.Intn(200)
	}
	fmt.Println(arr)
	ans := qSort(arr)
	fmt.Println(ans)
}
