package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println(p, *p)

	take(x)
	fmt.Println(x)
	
	takePonteiro(&x)
	fmt.Println(x)

	y := create()
	fmt.Println(*y)
}

// Passagem de valor para dentro da função
func take(x int) {
	fmt.Println(x)
	x = 100
}

// Passagem de referencia para dentro da função
func takePonteiro(x *int) {
	*x = 100
}

// Retorna um ponteiro para um inteiro
func create() *int {
	x := 10
	return &x
}
