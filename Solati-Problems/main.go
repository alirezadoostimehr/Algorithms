// O(n)
package main

import (
	"fmt"
)

func run(a []int) ([]int, []int) {
	cnt := make(map[int]int)

	for _, val := range a {
		cnt[val]++
	}
	un := make([]int, 0)
	nun := make([]int, 0)
	added := make(map[int]bool)
	for _, val := range a {
		if cnt[val] == 1 {
			un = append(un, val)
		} else if !added[val] {
			nun = append(nun, val)
		}
		added[val] = true
	}
	return un, nun
}

func main() {
	a := []int{1, 2, 6, 5, 4, 4}
	fmt.Print(run(a))
}
