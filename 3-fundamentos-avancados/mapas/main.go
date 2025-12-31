package main

import (
	"fmt"
)

func main() {
	var mapaNil map[string]string // mapa não inicializado
	fmt.Println(mapaNil == nil) // true

	mapaMake := make(map[string]string) // mapa inicializado
	fmt.Println(mapaMake == nil) // false	
	mapaMake["address"] = "Rua das pazes"
	fmt.Println(mapaMake)
	valor, hasKey := mapaMake["foo"]
	fmt.Println(valor, hasKey)

	mapaLiteral := map[string]string{
		"name": "Ariel",
		"role": "Tech Lead", 
	}
	fmt.Println(mapaLiteral)
	for k := range mapaLiteral {
		delete(mapaLiteral, k)
	}
	fmt.Println(mapaLiteral)

	mapaSlices := map[string][]string{
		"filmes": {"Avatar", "Chainsaw Man", "Jujutsu Kaizen"},
		"livros": {"Clean Code", "Iot Pratical Handbook"},
	}
	fmt.Println(mapaSlices)
	delete(mapaSlices, "filmes")
	fmt.Println(mapaSlices)
	clear(mapaSlices)
	fmt.Println(mapaSlices)
}
