// Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.
package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	numMin := 10000000000000000
	numMax := -1000000000000000

	for _, n := range slice {
		if n > numMax {
			numMax = n
		}
		if n < numMin {
			numMin = n
		}
	}
	fmt.Println("Valor mínimo:", numMin)
	fmt.Println("Valor máximo:", numMax)
}
