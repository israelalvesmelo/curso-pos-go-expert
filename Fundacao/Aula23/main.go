package main

func main() {
	for i := 0; i < 4; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for i, value := range numeros {
		println(i, value)
	}

	i := 0
	for i < 5 {
		println(i)
		i++
	}

	// for { //loop infinito
	// 	println(i)
	// }
}
