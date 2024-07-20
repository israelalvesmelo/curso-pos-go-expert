package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
	}
}

// Thread principal
func main() {
	// Second Thread
	go task("A")
	// Third Thread
	go task("B")
	// Fourth Thread
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(time.Second * 1)
		}
	}()

	time.Sleep(10 * time.Second)
}
