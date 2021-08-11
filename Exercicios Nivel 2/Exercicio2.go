package main

import (
	"fmt"
)
/*Exercicio de operadores*/
func main() {

	a := 10 == 9
	b := 100 != 10
	c := 5 <= 3
	d := 6 >= 3
	e := 8 < 10
	f := 7 > 3

	fmt.Printf("Verificação se há igualdade: %v\n", a)
	fmt.Printf("Verificação se é diferente: %v\n", b)
	fmt.Printf("Verificação se é menor ou igual : %v\n", c)
	fmt.Printf("Verificação se é maior ou igual : %v\n", d)
	fmt.Printf("Verificação se é menor : %v\n", e)
	fmt.Printf("Verificação se é maior : %v\n", f)
}
