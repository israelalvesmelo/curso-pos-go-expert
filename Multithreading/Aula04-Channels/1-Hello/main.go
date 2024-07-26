package main

import "time"

// Thread 1
func main() {
	channel := make(chan string) // Cria um canal de comunicação entre as threads
	// Thread 2
	go func() {
		time.Sleep(time.Second * 10)
		channel <- "Hello World" // Envia a mensagem para o canal
	}()

	//Thread 1
	msg := <-channel // Recebe a mensagem do canal
	println(msg)
}
