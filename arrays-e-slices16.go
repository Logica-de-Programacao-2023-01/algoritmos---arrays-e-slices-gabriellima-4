//Crie um Array de inteiros com 10 elementos.
//Crie um novo Slice que contenha apenas os elementos pares do Array original. Imprima o novo Slice.
package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceRes := []int{}

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			sliceRes = append(sliceRes, array[i])
		}
	}

	fmt.Println("Elementos pares do array original:", sliceRes)
}
