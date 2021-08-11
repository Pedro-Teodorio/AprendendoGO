/*Exercicio de deslocamento de Bits
Obs: Tambem dá para fazer com deslocamento de bits a multiplicação e a divisão por dois

<< -Multipicação
>> -Divisão

*/
package main

import (
	"fmt"
)

func main() {
	h := 895

	fmt.Printf("Valor decimal:%d\nValor Binário:%b\nValor HexaDecimal:%#x\n\n", h, h, h)

	j := h << 1

	fmt.Printf("Valor decimal:%d\nValor Binario:%b\nValor HexaDecimal:%#x", j, j, j)
}
