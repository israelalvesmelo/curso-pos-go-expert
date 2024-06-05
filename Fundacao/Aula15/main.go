package main

func main() {
	// Memória -> Endereço -> Valor
	// Ponteiro é um endereçamento da memoria
	a := 10                // Crio uma variavel
	println(a)             // Exibo o valor da variavel
	var ponteiro *int = &a // Crio uma variavel do tipo ponteiro que recebe o endereço da memoria da variaver "a"
	println(ponteiro)      // Exibo o endereçamento da memoria
	println(*ponteiro)     // Exibo o valor que está no endereço da memoria
	*ponteiro = 40         // Alterei o valor que está no endereço de memoria va variavel ponteiro
	println(*ponteiro)     // Exibo o valor que está no endereço da memoria

}
