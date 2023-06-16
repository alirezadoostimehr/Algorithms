// O(n)
package main

import (
	"fmt"
)

func run(a []int) int {
	res := 0
	cnt := 0
	for _, val := range a {
		if val > 0 {
			res += cnt
		} else {
			cnt++
		}
	}
	return res
}

func main() {
	a := []int{0, 1, 0, 1, 1}
	fmt.Print(run(a))
}
