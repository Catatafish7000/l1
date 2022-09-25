package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

const searchedVal int = 10

func bSearch(arr []int, val int) (int, bool) {
	var min, mid int
	max := len(arr) - 1

	for min <= max {
		mid = (min + max) / 2
		if arr[mid] == val {
			return mid, true
		}
		if arr[mid] < val {
			min = mid + 1
		}
		if arr[mid] > val {
			max = mid - 1
		}
	}
	return 0, false
}

func main() {
	sortedNums := make([]int, 100, 100)
	for i, _ := range sortedNums {
		sortedNums[i] = rand.Intn(100)
	}
	sort.Ints(sortedNums)
	fmt.Println(sortedNums)
	pos, ok := bSearch(sortedNums, searchedVal)
	if !ok {
		log.Println("Значение не найдено")
	} else {
		fmt.Printf("Значение %d находится на позиции %d\n", searchedVal, pos+1)
	}
}
