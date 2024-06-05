package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco //Endereco está compondo minha struct de client
}

func (c *Cliente) toString() {
	status := "ativo"
	if !c.Ativo {
		status = "inativo"
	}
	fmt.Printf("O nome do cliente é %s, sua idade é %d, ele(a) está %s", c.Nome, c.Idade, status)
}

func main() {
	cliente01 := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	cliente01.Logradouro = "Rua teste"
	cliente01.toString()
}
