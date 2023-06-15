package main

import "fmt"

func main() {
	var n int
	ans := 0
	fmt.Scan(n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(x)
		ans ^= x
	}
	fmt.Println(ans)
}
