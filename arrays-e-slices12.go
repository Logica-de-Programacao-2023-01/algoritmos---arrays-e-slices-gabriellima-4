// Crie um Array de inteiros com 5 elementos.
// Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.
package main

import "fmt"

func main() {
	array := [5]int{3, 7, 12, 24, 25}
	sliceResultante := []int{}

	for i := 0; i < len(array); i++ {
		if array[i]%3 == 0 {
			sliceResultante = append(sliceResultante, array[i])
		}
	}
	fmt.Println("Os múltiplos de 3 dentro do array são:", sliceResultante)
}
