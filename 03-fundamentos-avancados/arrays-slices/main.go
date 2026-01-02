package main

import "fmt"

func main() {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(arr) // imprimir array inteiro

	slice := arr[1:4] // [inclusivo:exclusivo]  indice 1 ao indice 3
	fmt.Println(slice)

	omitLimits := arr[:]
	fmt.Println(omitLimits) // imprimir array inteiro

	sliceLiteral := []int{1,2,3} // omitir o tamanho = slice
	fmt.Println(sliceLiteral, len(sliceLiteral), cap(sliceLiteral))

	var sliceNil []int // slice nullo
	fmt.Println(sliceNil == nil)

	filmesDB := [5]string{
		"All the bright places",
		"Black Mirror",
		"Inception",
		"Butterfly Effect",
		"Never Back Down",
	}

	filmes := make([]string, 0, 5) // slice com capacidade pre alocada

	for i := range 4 { // 1,2 ... 
		filme := filmesDB[i +1] 		
		filmes = append(filmes, filme)
	}
	fmt.Println(filmes)

	sliceCap := arr[:2:2] // full slice expression
	fmt.Println(sliceCap)

	_ = slice[2] // bounds check
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])

	slicePonteiro(slice)
	fmt.Println(slice[0])

	arrValor(arr)
	fmt.Println(arr[0])

	arrPonteiro(&arr)
	fmt.Println(arr[0])
}

func slicePonteiro(slice []int) {
	slice[0] = 200 // altera o valor no array original, pois um slice copia as referencias
}

func arrValor(arr [10]int) {
	arr[0] = 200 // não altera o valor no array original, pois copia valores
}

func arrPonteiro(arr *[10]int) {
	arr[0] = 200 // altera o valor no array original, pois explicitamente copia referencias
}
