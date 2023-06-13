// Crie um Array de floats com 10 elementos.
// Crie um novo Slice que contenha apenas os elementos do Array que são maiores que 5. Imprima o novo Slice.
package main

import "fmt"

func main() {
	array := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceRes := []float64{}

	for i := 0; i < len(array); i++ {
		if array[i] > 5 {
			sliceRes = append(sliceRes, array[i])
		}
	}
	fmt.Println("Elementos que são maiores do que 5:", sliceRes)

}
