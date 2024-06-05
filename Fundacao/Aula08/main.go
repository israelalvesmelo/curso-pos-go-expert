package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3))
}

func sum(a, b int) (int, bool) {
	sum := a + b
	if sum >= 50 {
		return sum, true
	}
	return sum, false
}
