/*Exercicio de Conversão de tipos*/
package main

import (
	"fmt"
)

type num int

var d num

var e int

func main() {

	fmt.Printf("%v\n%T\n", d, d)
	d = 42
	fmt.Println(d)

	e = int(d)

	fmt.Printf("%v,%T\n", e, e)
}
