package tipos

import (
	"fmt"
	"strconv"
)

func Tipos() {
	var isActive bool = true
	var age uint = 18
	var debit int = -10
	var height float32 = 1.75
	var complexNumber complex64 = 1 + 2i
	var name string = "Ariel"

	ageFloat := float64(age)

	s := strconv.FormatInt(int64(age), 10)

	fmt.Println(s) // "18"
}

func constantes() {
	const pi float64 = 3.14
	const e = 2.71828
	const name = "Ariel"

	fmt.Println(pi, e, name)
}