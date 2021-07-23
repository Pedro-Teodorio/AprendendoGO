/*Exercicio de Criação de tipo*/
package main

import (
	"fmt"
)

type num int

var x num

func main() {

	fmt.Printf("%v\n%T\n", x, x)

	x = 42

	fmt.Println(x)
}
