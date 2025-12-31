package main

import "fmt"

func main() {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(arr) // print full arr

	slice := arr[1:4] // [inclusivo:exclusivo]  indice 1 ao indice 3
	fmt.Println(slice)

	omitLimits := arr[:]
	fmt.Println(omitLimits) // imprimir array inteiro

	sliceLiteral := []int{1,2,3} // omitir o tamanho = slice
	fmt.Println(sliceLiteral, len(sliceLiteral), cap(sliceLiteral))

	var sliceNil []int // slice nullo
	fmt.Println(sliceNil == nil)
}

