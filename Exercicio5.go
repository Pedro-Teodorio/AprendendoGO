/*Exercicio de Conversão de tipos*/
package main

import (
	"fmt"
)

type num int

var x num

var y int

func main() {

	fmt.Printf("%v\n%T\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)

	fmt.Printf("%v,%T\n", y, y)
}
