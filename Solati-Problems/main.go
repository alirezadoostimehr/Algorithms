// O(n)
package main

import (
	"fmt"
	"sort"
)

func run(a []int, n int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a[len(a)-n]
}

func main() {
	a := []int{1, 2, 6, 5, 4, 4}
	fmt.Print(run(a, 1))
}
