// O(n)
package main

import (
	"fmt"
)

func run(a []int, n int) []int {
	n %= len(a)
	a = append(a[len(a)-n:], a[:len(a)-n]...)
	return a
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Print(run(a, 3))
}
