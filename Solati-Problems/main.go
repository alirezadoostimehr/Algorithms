// O(n)
package main

import (
	"fmt"
)

type pii struct {
	first  int
	second int
}

func run(a []int) []pii {
	res := make([]pii, 0)
	zeros := make([]int, 0)
	for i, val := range a {
		if val > 0 {
			for _, zero := range zeros {
				res = append(res, pii{zero, i})
			}
		} else {
			zeros = append(zeros, i)
		}
	}
	return res
}

func main() {
	a := []int{0, 1, 0, 1, 1}
	fmt.Print(run(a))
}
