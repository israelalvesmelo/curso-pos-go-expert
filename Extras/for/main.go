package main

import (
	"fmt"
)

func main() {
	x := 10
	for i := range x {
		fmt.Println(i)
	}

	done := make(chan string)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- fmt.Sprintf("Foi processado o valor: %s", v)

		}()
	}

	for range values {
		fmt.Println(<-done)
	}
}
