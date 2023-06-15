// O(n)
// It also can be done in O(n * Log(n)) with using map
package main

import "fmt"

func run(a []int) int {
	result := 0
	for _, val := range a {
		result ^= val
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 1, 2}
	fmt.Print(run(a))
}
