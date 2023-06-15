// O(n)
// It also can be done in O(n * Log(n)) with using map
package main

import (
	"fmt"
)

type pii struct {
	first  int
	secont int
}

func run(a []int, sum int) pii {
	ind := make(map[int]int)
	for i, val := range a {
		if ind[sum-val] != 0 {
			return pii{ind[sum-val] - 1, i}
		}
		ind[val] = i + 1
	}
	return pii{-1, -1}
}

func main() {
	a := []int{1, 2, 3, 5, 7, 11}
	fmt.Print(run(a, 16))
}
