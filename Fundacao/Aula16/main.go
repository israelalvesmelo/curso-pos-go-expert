package main

func soma(a, b *int) int { //O metodo recebe os endereços de memoria(ponteiros)
	*a = 50 // altero os valores que estão no endereço da memoria
	*b = 50
	return *a + *b // retorno a soma dos valores
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	soma(&minhaVar1, &minhaVar2) // Envio o endereço da memoria

	println(minhaVar1)
	println(minhaVar2)
}

// Ponteiros são uteis quando:
// - Quero reaproveitar o valor de uma variavel em outro scope
// - Não quero duplicar valores na memoria
// - Quero criar valores multaveis
