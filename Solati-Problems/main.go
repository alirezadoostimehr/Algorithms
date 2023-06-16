// O(n)
package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func run(a []int) int {
	sum := 0
	for _, val := range a {
		sum += val
	}

	res := abs(sum)
	prefixSum := 0
	suffixSum := sum
	for _, val := range a {
		prefixSum += val
		suffixSum -= val
		res = min(res, abs(suffixSum-prefixSum))
	}

	return res
}

func main() {
	a := []int{3, 1, 2, 4, 3}
	fmt.Print(run(a))
}
