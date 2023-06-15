// O(n)
package main

import (
	"fmt"
	"sort"
)

func run(a []int, n int) int {
	cnt := make(map[int]int)
	for _, val := range a {
		cnt[val]++
	}
	keys := make([]int, 0)
	for key := range cnt {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	n = len(a) - n
	now := 0
	for _, val := range keys {
		now += cnt[val]
		if now >= n {
			return val
		}

	}
	return 0
}

func main() {
	a := []int{1, 2, 6, 5, 4, 4}
	fmt.Print(run(a, 3))
}
