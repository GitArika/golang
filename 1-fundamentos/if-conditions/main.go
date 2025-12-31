package main

import (
	"fmt"
	"errors"
)

func main() {
	ifVariavelEscopo()
	ifBool()
}

func ifVariavelEscopo() {
	if x := 5; x < 0 {
		fmt.Println(" menor que 0 ")
	} else if x >= 10 {
		fmt.Println(" maior ou igual a 10")
	} else {
		fmt.Println(" edge case ")
	}
}

func doError() error {
	return errors.New("error")
}

func ifBool() {
	err := doError()

	if err != nil {
		fmt.Println(" ocorreu um erro ")
	}
}
