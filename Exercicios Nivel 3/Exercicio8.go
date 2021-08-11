package main

import (
	"fmt"
)

func main() {
	cansaso := 100

	switch {
	case cansaso == 0:
		{
			fmt.Println("Você está pronto para Tudo")
		}
	case cansaso == 50:
		{
			fmt.Println("Você ainda aguenta muita coisa")
		}
	case cansaso == 100:
		{
			fmt.Println("Mano vai Dormir")
		}
	}
}
