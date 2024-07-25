package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int) {
	for v := range ch {
		fmt.Printf("Received %d\n", v)
	}
}
