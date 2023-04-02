// Crie um Array de inteiros com 10 elementos.
// Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array.
// Imprima o resultado da busca.
package main

import "fmt"

func main() {
	var x int
	var y int
	var numeros = [10]int{1, 23, 7, 44, 5, 38, 0, 68, 91, 100}
	fmt.Print("Insira o valor que deseja procurar.")
	fmt.Scan(&x)

	for {
		y++
		if x == numeros[y] {
			fmt.Println("O valor que você procura já está inserido no array.")
		} else if len(numeros) >= 10 {
			fmt.Println("O valor que você procura não está presente no array.")
			break
		}
	}
}
