package main

import "fmt"

func foo (a interface{}) {} // interface{} == any

type Animal interface {
	Sound() string
}

type Dog struct{}

func (Dog) Sound() string {
	return "Auuuuuuuuuuuuu"
}

func (Dog) String() string {
	return "Sou um Dog"
}

type Cat struct{}

func (Cat) Sound() string {
	return "Miauuuuuuu"
}

func (Cat) String() string {
	return "Sou um Cat"
}


func makeSound(a Animal) {
	fmt.Println(a.Sound())
}

func takeEmptyInterface(a any) {
	str, ok := a.(string) // type assertion
	fmt.Println(str, ok)
}

func takeAnimal(a Animal) {
	switch t := a.(type) {
		case Dog, Cat:
			fmt.Println(t.Sound())
		default:
			fmt.Println("unknown type")
	}
}

func main() {
	dog := Dog{}
	makeSound(dog)
	takeEmptyInterface(10)
	takeEmptyInterface("b")
	takeEmptyInterface(dog)
	var a Animal
	takeEmptyInterface(a)

	takeAnimal(dog)
	takeAnimal(a)
	cat := Cat{}
	takeAnimal(cat)

	fmt.Println(dog)
	fmt.Println(cat)
}

