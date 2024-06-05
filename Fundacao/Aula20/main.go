package main

// Isso é uma generic
func soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber float32

type Number interface { //Isso se chama constrains
	int | ~float32
}

func somaGeneric2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func compara[T comparable](a T, b T) bool { // Comparable serve para comparar valores do mesmo tipo
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "Joao": 3000, "Maria": 400}
	m2 := map[string]float64{"Wesley": 30.0, "Joao": 10.30, "Maria": 400}
	m3 := map[string]MyNumber{"Wesley": 30.0, "Joao": 10.30, "Maria": 400}

	println(soma(m))
	println(soma(m2))
	println(somaGeneric2(m3))
	println(compara(10, 30))
}
