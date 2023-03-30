// Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas.
// Solicite ao usuário que informe os valores de cada elemento da matriz. Em seguida, imprima a matriz resultante.
package main

import "fmt"

func main() {
	var matriz [3][2]int
	fmt.Println(matriz)

	var v1 int
	var v2 int
	var v3 int
	var v4 int
	var v5 int
	var v6 int

	fmt.Print("Escolha um número.")
	fmt.Scan(&v1)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v2)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v3)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v4)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v5)
	fmt.Print("Escolha um número.")
	fmt.Scan(&v6)

	matrizinicializada := [3][2]int{{v1, v2}, {v3, v4}, {v5, v6}}
	fmt.Println(matrizinicializada)
	
}
