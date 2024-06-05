package main

import "fmt"

const wording = "Hello, World"

type ID int

var (
	f ID = 1
)

func main() {
	fmt.Printf("O tipo de F é %T. \n", f)
	fmt.Printf("O valor de F é %v", f)

}
