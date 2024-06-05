package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/israelalvesmelo/curso-pos-go-expert/Fundacao/Aula21/matematica"
)

func main() {
	s := matematica.Soma(1, 2)
	fmt.Println("O valor da soma é ", s)

	fmt.Println(uuid.New())
}
