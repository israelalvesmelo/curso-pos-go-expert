package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
}

type Pessoa interface {
	desativar()
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco //Endereco está compondo minha struct de client
}

func (c *Cliente) desativar() {
	c.Ativo = false
}

func (c *Cliente) toString() {
	status := "ativo"
	if !c.Ativo {
		status = "inativo"
	}
	fmt.Printf("O nome do cliente é %s, sua idade é %d, ele(a) está %s.", c.Nome, c.Idade, status)
}

func desativacao(pessoa Pessoa) {
	pessoa.desativar()
}

func main() {
	cliente01 := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	desativacao(&cliente01)
	cliente01.toString()
}
