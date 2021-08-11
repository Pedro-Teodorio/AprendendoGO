package main

import (
	"fmt"
)

const (
	ano = (iota + 2021)
	ano1
	ano2
	ano3
	ano4
)

func main() {

	fmt.Printf("Proximos 4 anos: %v,%v,%v,%v", ano1, ano2, ano3, ano4)

}
