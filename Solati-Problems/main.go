// O(n)
package main

import (
	"fmt"
)

func run(a []int, n int) int {
	cnt := make(map[int]int)
	for _, val := range a {
		cnt[val]++
	}
	n = len(a) - n
	now := 0
	for key, val := range cnt {
		if now+val >= n {
			return key
		}
		now += n

	}
	return 0
}

func main() {
	a := []int{1, 2, 6, 5, 4, 4}
	fmt.Print(run(a, 1))
}
