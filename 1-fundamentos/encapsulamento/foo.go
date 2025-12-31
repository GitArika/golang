package encapsulamento

import (
	"fmt"
	"my-first-go-project/encapsulamento/internal"
)

var Foo string = "Foo"

func PrintInternal() {
	var nome, sobrenome, genero = "Ariel", "Leão", 1

	var (
		cargo = "Programador"
		salario = 1000
	)

	idade := 20

	fmt.Println(nome, sobrenome, genero, idade, cargo, salario)

	fmt.Println(internal.Interno)

	fmt.Println(somar(1, 2))

	a, b := swap(10, 20)
	fmt.Println(a, b)

	res, rem := dividir(10, 3)
	fmt.Println(res, rem)	

	x := somarHighOrder(10)
	fmt.Println(x(2))

	fmt.Println(somarVariatica(10, 10, 10))
}

func somar(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a,b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

func somarHighOrder(a int) func (int) int {
	return func (b int) int {
		return a + b
	}
}

func somarVariatica(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}