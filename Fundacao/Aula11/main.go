package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
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

	cliente01.toString()
}
