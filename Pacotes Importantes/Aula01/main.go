package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo \n"))
	//tamanho, err := f.WriteString("Escrevendo alguma string"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamnho: %d bytes \n", tamanho)
	f.Close()

	// leitura
	arquivo, err := os.ReadFile("arquivo.txt") //Aqui eu faço a abertura e a leitura do arquivo, indicado para arquivos pequeno
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo)) //Aqui estamos convertendo um slice de bytes em string

	// leitura linha a linha
	arquivo2, err := os.Open("arquivo.txt") //Aqui estou abrindo um arquivo
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2) // Crio um reader responsavel por percorrer o conteudo do arquivo aberto
	buffer := make([]byte, 10)          // O buffer é responsavel por definir o tamnho do byte que iremos ler
	for {
		indice, err := reader.Read(buffer) // com base no tamanho do buffer, inserimos o conteudo do reader nele
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:indice])) // Fazemos uma conversão do byte para string, além disso é exibido o conteudo que está no buffer usando o indice do for
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
