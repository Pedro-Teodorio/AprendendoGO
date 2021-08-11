package main

import (
	"fmt"
)

func main() {
	ano_de_nascimento := 2001
	ano_atual := 2021

	for {
		if ano_de_nascimento > ano_atual {
			break
		}
		fmt.Println(ano_de_nascimento)
		ano_de_nascimento++
	}
}
