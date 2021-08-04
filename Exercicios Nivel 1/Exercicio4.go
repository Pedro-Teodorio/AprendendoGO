/*Exercicio de Criação de tipo*/
package main

import (
	"fmt"
)

type number int

var p number

func main() {

	fmt.Printf("%v\n%T\n", p, p)

	p = 42

	fmt.Println(p)
}
