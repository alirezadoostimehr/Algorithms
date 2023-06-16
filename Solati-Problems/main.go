// O(n)
package main

import (
	"fmt"
)

func run(a []int) bool {
	exists := make(map[int]bool)
	sum := 0
	for _, val := range a {
		if exists[val] {
			return false
		}
		exists[val] = true
		sum += val
	}
	n := len(a)
	return sum == (n*(n+1))/2
}

func main() {
	a := []int{1, 4}
	fmt.Print(run(a))
}
