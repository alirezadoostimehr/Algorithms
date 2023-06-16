// O(n)
package main

import (
	"fmt"
	"math"
)

func run(a []int) int {
	sum := 0
	for _, val := range a {
		sum += val
	}
	res := (int)(math.Abs((float64)(sum)))
	s1 := 0
	s2 := sum
	for _, val := range a {
		s1 += val
		s2 -= val
		res = (int)(math.Min(math.Abs((float64)(s1-s2)), (float64)(res)))
	}
	return res
}

func main() {
	a := []int{3, 1, 2, 4, 3}
	fmt.Print(run(a))
}
