package main

import "fmt"

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Olá mundo", hello)
	ler(hello)
}

// receive only
func recebe(nome string, hello chan<- string) { // a seta ao lado direito do chan (<-) significa que o canal só recebe dados
	hello <- nome
}

// send only
func ler(data <-chan string) { // a seta ao lado esquerdo do chan (<-) significa que o canal só le dados
	fmt.Println(<-data)
}
