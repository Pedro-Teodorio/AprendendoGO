/*Exercicio de Entendimento de constantes tipadas e não tipadas*/
package main

import (
	"fmt"
)

const x = 100.0
const o = "Vamos Aprender Go"

func main() {
	fmt.Printf("Constante não tipada: %v, Tipo:%T\nConstante tipada: %v, Tipo:%T", x, x, o, o)
}
