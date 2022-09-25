package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	tempRec := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Float64s(tempRec)
	fmt.Println(tempRec)
	left := tempRec[0]
	right := tempRec[len(tempRec)-1]
	tMin := int(left) / 10 * 10
	tMax := int(right) / 10 * 10
	fmt.Println(tMin, tMax)
	groups := make(map[int][]float64)
Loop:
	for _, val := range tempRec {
		for temp := tMin; temp <= tMax; temp += 10 {
			if math.Abs(val) > math.Abs(float64(temp)) && math.Abs(val) < math.Abs(float64(temp))+10 && math.Signbit(val) == math.Signbit(float64(temp)) {
				groups[temp] = append(groups[temp], val)
				continue Loop
			}
		}
	}
	fmt.Println(groups)
}
