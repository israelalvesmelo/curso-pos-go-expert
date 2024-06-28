package main

import (
	"fmt"

	"github.com/israelalvesmelo/curso-pos-go-expert/Packaging/aula01/math"
)

func main() {
	var m math.Math
	m.A = 1
	m.B = 2
	fmt.Printf("A + B = %d\n", m.Add())
}
