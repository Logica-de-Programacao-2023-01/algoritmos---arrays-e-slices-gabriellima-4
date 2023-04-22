// Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor.
// Remova todas as ocorrências desse valor no Slice e imprima o resultado.
package main

import "fmt"

func main() {
	var numeros = []string{"Marta", "Daniel", "Gabriel", "Artur", "Luiza", "Gabriel", "Ana", "Beatriz"}
	var valor string
	var x int

	fmt.Print("Escreva um nome.")
	fmt.Scan(&valor)

	w := x + 1

	for len(numeros) > x {
		if valor == numeros[x] {
			numeros = append(numeros[:x], numeros[w:]...)
		}
		w++
		x++
	}

}
