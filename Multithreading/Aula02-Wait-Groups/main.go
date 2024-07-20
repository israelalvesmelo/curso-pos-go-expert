package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}

// Thread principal
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)
	// Second Thread
	go task("A", &waitGroup)
	// Third Thread
	go task("B", &waitGroup)
	// Fourth Thread
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(time.Second * 1)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
