package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "Carlos": 4000}
	fmt.Println(salarios)
	delete(salarios, "Wesley")
	salarios["Paula"] = 5000
	salarios["Sergio"] = 5000
	fmt.Printf("%v\n", salarios)

	/*
		Outras formas de criar

		sal := make(map[string]int)
		sal := map[string]int{}
		sal["Wesley"] = 1000
	*/

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

}
