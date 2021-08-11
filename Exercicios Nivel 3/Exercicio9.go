package main

import (
	"fmt"
)

func main() {

	esporte_favorito := "Basquete"

	switch esporte_favorito {
	case "Futebol":
		fmt.Println("Você é normal")
	case "Basquete":
		fmt.Println("Nossa bora ser Amigos")
	case "Volei":
		fmt.Println("Você gente boa")
	}

}
