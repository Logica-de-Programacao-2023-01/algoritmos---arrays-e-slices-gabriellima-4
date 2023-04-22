// Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas.
// Em seguida, solicite ao usuário que informe um índice de linha e outro de coluna e imprima o valor armazenado nessa posição da matriz.
package main

import "fmt"

func main() {
	var matriz [2][3]int
	fmt.Println(matriz)

	matrizinicializada := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrizinicializada)

	var indxLinha int
	var indxColuna int

	fmt.Print("Escolha um número.")
	fmt.Scan(&indxLinha)
	fmt.Print("Escolha um número.")
	fmt.Scan(&indxColuna)

	if indxLinha == 0 && indxColuna == 0 {
		fmt.Println(matrizinicializada[0][0])
	} else if indxLinha == 0 && indxColuna == 1 {
		fmt.Println(matrizinicializada[0][1])
	} else if indxLinha == 0 && indxColuna == 2 {
		fmt.Println(matrizinicializada[0][2])
	} else if indxLinha == 1 && indxColuna == 0 {
		fmt.Println(matrizinicializada[1][0])
	} else if indxLinha == 1 && indxColuna == 1 {
		fmt.Println(matrizinicializada[1][1])
	} else if indxLinha == 1 && indxColuna == 2 {
		fmt.Println(matrizinicializada[1][2])
	} else {
		fmt.Println("Índices não existem nessa matriz, portanto não há valor.")
	}
}
