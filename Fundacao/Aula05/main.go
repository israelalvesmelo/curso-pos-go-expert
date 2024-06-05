package main

import "fmt"

const wording = "Hello, World"

type ID int

var (
	f ID = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 22
	meuArray[2] = 33

	fmt.Println("O tamanho do array é: ", len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é: %d\n", i, v)
	}

}
