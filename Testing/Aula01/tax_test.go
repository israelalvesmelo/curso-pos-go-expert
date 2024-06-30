package aula01

import "testing"

// Para executar os testes, execute o comando abaixo:
// go test -v para executar todos os testes ou go test -v -run TestCalculateTax para executar apenas o TestCalculateTax.
// go test -coverprofile=coverage para gerar um relatório de cobertura de código.
// go tool cover -html=coverage para gerar um relatório de cobertura de código em HTML.

func TestCalculateTax(t *testing.T) {
	expected := 5.0
	result := CalculateTax(500.0)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calculateTax struct {
		amount   float64
		expected float64
	}
	table := []calculateTax{
		{1000.0, 10.0},
		{500.0, 5.0},
		{1500.0, 10.0},
	}

	for _, row := range table {
		result := CalculateTax(row.amount)
		if result != row.expected {
			t.Errorf("Expected %f, got %f", row.expected, result)
		}
	}
}

/*
	Benchmarking, no contexto de programação e engenharia de software, é o processo de medir o desempenho de um componente específico de um sistema,
	como uma função ou um bloco de código. O objetivo do benchmarking é determinar a eficiência de uma determinada parte do código em termos de tempo de execução,
 	uso de memória ou outros recursos.

	Para executar:
		go test -bench=.
	Executar um benchmark específico:
		go test -bench=BenchmarkAdd
	Executar todos os benchmarks e ignorar testes:
		go test -bench=. -run=^$
*/
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000.0)
	}
}

/*
	Fuzzing, também conhecido como "fuzz testing", é uma técnica de teste de software que envolve fornecer dados de entrada inválidos, inesperados ou aleatórios para um programa.
	O objetivo do fuzzing é descobrir bugs, falhas de segurança e outras vulnerabilidades ao ver como o programa lida com esses dados anômalos.
	Fuzzing é particularmente eficaz para encontrar problemas de buffer overflow, injeções de código e outras vulnerabilidades que podem ser exploradas por atacantes.

	Para executar:
		go test -fuzz=.
	Executar um fuzz específico:
		go test -fuzz=FuzzCalculateTax
	Executar todos os fuzz e ignorar testes:
		go test -fuzz=. -run=^$

		go test -fuzz=. -fuzztime=5s para executar fuzzing por 5 segundos.
*/
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0} // Seed with some values to start with and avoid a few false positives.
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})
}
