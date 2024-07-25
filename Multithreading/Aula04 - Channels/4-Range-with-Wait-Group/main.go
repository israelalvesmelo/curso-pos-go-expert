package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(ch)
	go reader(ch, &wg)
	wg.Wait()
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Printf("Received %d\n", v)
		wg.Done()
	}
}
