// O(n)
package main

import (
	"fmt"
)

func run(a []int) int {
	n := len(a) + 1
	sum := (n * (n + 1)) / 2
	for _, val := range a {
		sum -= val
	}
	return sum
}

func main() {
	a := []int{1, 2, 6, 5, 4}
	fmt.Print(run(a))
}
