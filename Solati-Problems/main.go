// O(n)
package main

import (
	"fmt"
)

func run(a []int, n int) []int {
	res := make([]int, 0)
	for i := 0; i < len(a); i++ {
		res = append(res, a[((i-n)+len(a))%len(a)])
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Print(run(a, 3))
}
