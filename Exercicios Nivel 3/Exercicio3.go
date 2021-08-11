package main

import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ {
		fmt.Print(i, "\n")

		for rep := 1; rep <= 3; rep++ {
			fmt.Printf("\t%#U\n", i)
		}
		fmt.Print("\n")

	}

}
